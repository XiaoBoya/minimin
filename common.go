package minimin

type Status string

const (
	Success Status = "SUCCESS" // 成功
	Fail    Status = "FAIL"    // 失败
	Run     Status = "RUNNING" // 运行中
	Init    Status = "INIT"    // 初始化
	Queue   Status = "QUEUE"   // 队列中

	SimpleFilePerm = 0666 // 普通文件权限
	SimpleDirPerm  = 0755 // 普通文件夹权限
	DangerPerm     = 0777 // 危险权限

	AppListFile    = "apps.json" // app列表存储文件名称
	EnvJsonFile    = ".env.json"
	ConfigYamlFile = "config.yaml"  // yaml配置文件名
	ConfigJsonFile = "config.json"  // json配置文件夹名
	InfoDir        = "info"         // 信息文件夹名
	RunLogJsonFile = "run_log.json" // json配置文件夹名
)
