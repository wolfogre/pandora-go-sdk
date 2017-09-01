package base

const (
	MethodGet    string = "GET"
	MethodPost   string = "POST"
	MethodDelete string = "DELETE"
	MethodPut    string = "PUT"
)

const (
	HTTPHeaderAppId         string = "X-AppId"
	HTTPHeaderContentType   string = "Content-Type"
	HTTPHeaderContentLength string = "Content-Length"
	HTTPHeaderContentMD5    string = "Content-MD5"
	HTTPHeaderRequestId     string = "X-Reqid"
	HTTPHeaderAuthorization string = "Authorization"
)

const (
	OpCreateGroup         string = "CreateGroup"
	OpUpdateGroup         string = "UpdateGroup"
	OpStartGroupTask      string = "StartGroupTask"
	OpStopGroupTask       string = "StopGroupTask"
	OpListGroups          string = "ListGroup"
	OpGetGroup            string = "GetGroup"
	OpDeleteGroup         string = "DeleteGroup"
	OpCreateRepo          string = "CreateRepo"
	OpGetRepo             string = "GetRepo"
	OpListRepos           string = "ListRepo"
	OpDeleteRepo          string = "DeleteRepo"
	OpPostData            string = "PostData"
	OpCreateTransform     string = "CreateTransform"
	OpUpdateTransform     string = "UpdateTransform"
	OpGetTransform        string = "GetTransform"
	OpListTransforms      string = "ListTransform"
	OpDeleteTransform     string = "DeleteTransform"
	OpCreateExport        string = "CreateExport"
	OpUpdateExport        string = "UpdateExport"
	OpGetExport           string = "GetExport"
	OpListExports         string = "ListExport"
	OpDeleteExport        string = "DeleteExport"
	OpUploadPlugin        string = "UploadPlugin"
	OpGetPlugin           string = "GetPlugin"
	OpListPlugins         string = "ListPlugin"
	OpDeletePlugin        string = "DeletePlugin"
	OpCreateDatasource    string = "CreateDatasource"
	OpGetDatasource       string = "GetDatasource"
	OpListDatasources     string = "ListDatasources"
	OpDeleteDatasource    string = "DeleteDatasource"
	OpCreateJob           string = "CreateJob"
	OpGetJob              string = "GetJob"
	OpListJobs            string = "ListJobs"
	OpDeleteJob           string = "DeleteJob"
	OpStartJob            string = "StartJob"
	OpStopJob             string = "StopJob"
	OpGetJobHistory       string = "GetJobHistory"
	OpStopJobBatch        string = "StopJobBatch"
	OpRerunJobBatch       string = "RerunJobBatch"
	OpCreateJobExport     string = "CreateJobExport"
	OpGetJobExport        string = "GetJobExport"
	OpListJobExports      string = "ListJobExports"
	OpDeleteJobExport     string = "DeleteJobExport"
	OpRetrieveSchema      string = "RetrieveSchema"
	OpUploadUdf           string = "UploadUdf"
	OpPutUdfMeta          string = "PutUdfMeta"
	OpDeleteUdf           string = "DeleteUdf"
	OpListUdfs            string = "ListUdfs"
	OpRegUdfFunc          string = "RegisterUdfFunction"
	OpDeregUdfFunc        string = "DeregisterUdfFunction"
	OpListUdfFuncs        string = "ListUdfFuncs"
	OpListUdfBuiltinFuncs string = "ListUdfBuiltinFuncs"

	OpUpdateRepo        string = "UpdateRepo"
	OpSendLog           string = "SendLog"
	OpQueryLog          string = "QueryLog"
	OpQueryHistogramLog string = "QueryHistogramLog"
	OpPutRepoConfig     string = "PutRepoConfig"
	OpGetRepoConfig     string = "GetRepoConfig"
	OpPartialQuery      string = "PartialQuery"

	OpUpdateRepoMetadata string = "UpdataRepoMetadata"
	OpDeleteRepoMetadata string = "DeleteRepoMetadata"
	OpUpdateViewMetadata string = "UpdataViewMetadata"
	OpDeleteViewMetadata string = "DeleteViewMetadata"

	OpCreateSeries         string = "CreateSeries"
	OpUpdateSeriesMetadata string = "UpdataSeriesMetadata"
	OpDeleteSeriesMetadata string = "DeleteSeriesMetadata"
	OpListSeries           string = "ListSeries"
	OpDeleteSeries         string = "DeleteSeries"

	OpCreateView  string = "CreateView"
	OpListView    string = "ListView"
	OpGetView     string = "GetView"
	OpDeleteView  string = "DeleteView"
	OpQueryPoints string = "QueryPoints"
	OpWritePoints string = "WritePoints"

	OpActivateUser   string = "ActivateUser"
	OpCreateDatabase string = "CreateDatabase"
	OpListDatabases  string = "ListDatabases"
	OpDeleteDatabase string = "DeleteDatabase"
	OpCreateTable    string = "CreateTable"
	OpUpdateTable    string = "UpdateTable"
	OpListTables     string = "ListTables"
	OpDeleteTable    string = "DeleteTable"
	OpGetTable       string = "GetTable"
)

const (
	ContentTypeJson        string = "application/json"
	ContentTypeJar         string = "application/java-archive"
	ContentTypeText        string = "text/plain"
	ContentTypeOctetStream string = "application/octet-stream"
)

const (
	NestLimit int = 5
)
