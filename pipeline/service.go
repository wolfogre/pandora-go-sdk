package pipeline

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/qiniu/pandora-go-sdk/base"
	"github.com/qiniu/pandora-go-sdk/base/config"
	"github.com/qiniu/pandora-go-sdk/base/ratelimit"
	"github.com/qiniu/pandora-go-sdk/base/request"
	"github.com/qiniu/pandora-go-sdk/logdb"
	"github.com/qiniu/pandora-go-sdk/tsdb"
)

var builder errBuilder

type Pipeline struct {
	Config        *config.Config
	HTTPClient    *http.Client
	reqLimit      *ratelimit.Limiter
	flowLimit     *ratelimit.Limiter
	repoSchemas   map[string]RepoSchema
	repoSchemaMux sync.Mutex
	defaultRegion string

	LogDB logdb.LogdbAPI
	TSDB  tsdb.TsdbAPI
}

type RepoSchema map[string]RepoSchemaEntry

func NewConfig() *config.Config {
	return config.NewConfig()
}

func New(c *config.Config) (PipelineAPI, error) {
	return newClient(c)
}

func (c *Pipeline) Close() (err error) {
	if c.reqLimit != nil {
		err = c.reqLimit.Close()
		if err != nil {
			c.Config.Logger.Errorf("Close reqLimit error %v", err)
		}
	}
	if c.flowLimit != nil {
		err = c.flowLimit.Close()
		if err != nil {
			c.Config.Logger.Errorf("Close flowLimit error %v", err)
		}
	}
	return
}

func newClient(c *config.Config) (p *Pipeline, err error) {
	if c.PipelineEndpoint == "" {
		c.PipelineEndpoint = c.Endpoint
	}
	if c.PipelineEndpoint == "" {
		c.PipelineEndpoint = config.DefaultPipelineEndpoint
	}
	c.ConfigType = config.TypePipeline
	if err = base.CheckEndPoint(c.PipelineEndpoint); err != nil {
		return
	}
	var t = &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   c.DialTimeout,
			KeepAlive: 30 * time.Second,
		}).Dial,
		ResponseHeaderTimeout: c.ResponseTimeout,
	}

	p = &Pipeline{
		Config:        c,
		HTTPClient:    &http.Client{Transport: t},
		repoSchemas:   make(map[string]RepoSchema),
		repoSchemaMux: sync.Mutex{},
		defaultRegion: defaultRegion,
	}
	if c.RequestRateLimit > 0 {
		p.reqLimit = ratelimit.NewLimiter(c.RequestRateLimit)
	}
	if c.FlowRateLimit > 0 {
		p.flowLimit = ratelimit.NewLimiter(1024 * c.FlowRateLimit)
	}
	return
}

func (c *Pipeline) newRequest(op *request.Operation, token string, v interface{}) *request.Request {
	req := request.New(c.Config, c.HTTPClient, op, token, builder, v)
	req.Data = v
	return req
}

func (c *Pipeline) newOperation(opName string, args ...interface{}) *request.Operation {
	var method, urlTmpl string
	switch opName {
	case base.OpCreateGroup:
		method, urlTmpl = base.MethodPost, "/v2/groups/%s"
	case base.OpUpdateGroup:
		method, urlTmpl = base.MethodPut, "/v2/groups/%s"
	case base.OpStartGroupTask:
		method, urlTmpl = base.MethodPost, "/v2/groups/%s/actions/start"
	case base.OpStopGroupTask:
		method, urlTmpl = base.MethodPost, "/v2/groups/%s/actions/stop"
	case base.OpListGroups:
		method, urlTmpl = base.MethodGet, "/v2/groups"
	case base.OpGetGroup:
		method, urlTmpl = base.MethodGet, "/v2/groups/%s"
	case base.OpDeleteGroup:
		method, urlTmpl = base.MethodDelete, "/v2/groups/%s"
	case base.OpCreateRepo:
		method, urlTmpl = base.MethodPost, "/v2/repos/%s"
	case base.OpUpdateRepo:
		method, urlTmpl = base.MethodPut, "/v2/repos/%s"
	case base.OpListRepos:
		method, urlTmpl = base.MethodGet, "/v2/repos"
	case base.OpGetRepo:
		method, urlTmpl = base.MethodGet, "/v2/repos/%s"
	case base.OpDeleteRepo:
		method, urlTmpl = base.MethodDelete, "/v2/repos/%s"
	case base.OpPostData:
		method, urlTmpl = base.MethodPost, "/v2/repos/%s/data"
	case base.OpCreateTransform:
		method, urlTmpl = base.MethodPost, "/v2/repos/%s/transforms/%s/to/%s"
	case base.OpUpdateTransform:
		method, urlTmpl = base.MethodPut, "/v2/repos/%s/transforms/%s"
	case base.OpListTransforms:
		method, urlTmpl = base.MethodGet, "/v2/repos/%s/transforms"
	case base.OpGetTransform:
		method, urlTmpl = base.MethodGet, "/v2/repos/%s/transforms/%s"
	case base.OpDeleteTransform:
		method, urlTmpl = base.MethodDelete, "/v2/repos/%s/transforms/%s"
	case base.OpCreateExport:
		method, urlTmpl = base.MethodPost, "/v2/repos/%s/exports/%s"
	case base.OpUpdateExport:
		method, urlTmpl = base.MethodPut, "/v2/repos/%s/exports/%s"
	case base.OpListExports:
		method, urlTmpl = base.MethodGet, "/v2/repos/%s/exports"
	case base.OpGetExport:
		method, urlTmpl = base.MethodGet, "/v2/repos/%s/exports/%s"
	case base.OpDeleteExport:
		method, urlTmpl = base.MethodDelete, "/v2/repos/%s/exports/%s"
	case base.OpUploadPlugin:
		method, urlTmpl = base.MethodPost, "/v2/plugins/%s"
	case base.OpListPlugins:
		method, urlTmpl = base.MethodGet, "/v2/plugins"
	case base.OpGetPlugin:
		method, urlTmpl = base.MethodGet, "/v2/plugins/%s"
	case base.OpDeletePlugin:
		method, urlTmpl = base.MethodDelete, "/v2/plugins/%s"
	case base.OpCreateDatasource:
		method, urlTmpl = base.MethodPost, "/v2/datasources/%s"
	case base.OpGetDatasource:
		method, urlTmpl = base.MethodGet, "/v2/datasources/%s"
	case base.OpListDatasources:
		method, urlTmpl = base.MethodGet, "/v2/datasources"
	case base.OpDeleteDatasource:
		method, urlTmpl = base.MethodDelete, "/v2/datasources/%s"
	case base.OpCreateJob:
		method, urlTmpl = base.MethodPost, "/v2/jobs/%s"
	case base.OpGetJob:
		method, urlTmpl = base.MethodGet, "/v2/jobs/%s"
	case base.OpListJobs:
		method, urlTmpl = base.MethodGet, "/v2/jobs%s"
	case base.OpDeleteJob:
		method, urlTmpl = base.MethodDelete, "/v2/jobs/%s"
	case base.OpStartJob:
		method, urlTmpl = base.MethodPost, "/v2/jobs/%s/actions/start"
	case base.OpStopJob:
		method, urlTmpl = base.MethodPost, "/v2/jobs/%s/actions/stop"
	case base.OpGetJobHistory:
		method, urlTmpl = base.MethodGet, "/v2/jobs/%s/history"
	case base.OpStopJobBatch:
		method, urlTmpl = base.MethodPost, "/v2/batch/actions/stop"
	case base.OpRerunJobBatch:
		method, urlTmpl = base.MethodPost, "/v2/batch/actions/rerun"
	case base.OpCreateJobExport:
		method, urlTmpl = base.MethodPost, "/v2/jobs/%s/exports/%s"
	case base.OpGetJobExport:
		method, urlTmpl = base.MethodGet, "/v2/jobs/%s/exports/%s"
	case base.OpListJobExports:
		method, urlTmpl = base.MethodGet, "/v2/jobs/%s/exports"
	case base.OpDeleteJobExport:
		method, urlTmpl = base.MethodDelete, "/v2/jobs/%s/exports/%s"
	case base.OpRetrieveSchema:
		method, urlTmpl = base.MethodPost, "/v2/schemas"
	case base.OpUploadUdf:
		method, urlTmpl = base.MethodPost, "/v2/udf/jars/%s"
	case base.OpPutUdfInfo:
		method, urlTmpl = base.MethodPut, "/v2/udf/jars/%s"
	case base.OpDeleteUdf:
		method, urlTmpl = base.MethodDelete, "/v2/udf/jars/%s"
	case base.OpListUdf:
		method, urlTmpl = base.MethodGet, "/v2/udf/jars%s"
	case base.OpRegUdfFunc:
		method, urlTmpl = base.MethodPost, "/v2/udf/funcs/%s"
	case base.OpDeregUdfFunc:
		method, urlTmpl = base.MethodDelete, "/v2/udf/funcs/%s"
	case base.OpListUdfFuncs:
		method, urlTmpl = base.MethodGet, "/v2/udf/funcs%s"

	default:
		c.Config.Logger.Errorf("unmatched operation name: %s", opName)
		return nil
	}

	return &request.Operation{
		Name:   opName,
		Method: method,
		Path:   fmt.Sprintf(urlTmpl, args...),
	}
}

func (c *Pipeline) GetLogDBAPI() (logdb.LogdbAPI, error) {
	if c.LogDB == nil {
		logdb, err := logdb.New(c.Config.Clone())
		if err != nil {
			return nil, err
		}
		c.LogDB = logdb
	}
	return c.LogDB, nil
}

func (c *Pipeline) GetTSDBAPI() (tsdb.TsdbAPI, error) {
	if c.TSDB == nil {
		tsdb, err := tsdb.New(c.Config.Clone())
		if err != nil {
			return nil, err
		}
		c.TSDB = tsdb
	}
	return c.TSDB, nil
}
