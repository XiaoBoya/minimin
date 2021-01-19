package minimin

const (
	Success = "SUCCESS" // 成功
	Fail    = "FAIL"    // 失败
	Run     = "RUNNING" // 运行中
	Init    = "INIT"    // 初始化
	Queue   = "QUEUE"   // 队列中

	SimpleFilePerm = 0666 // 普通文件权限
	SimpleDirPerm  = 0755 // 普通文件夹权限
	DangerPerm     = 0777 // 危险权限

	InfoDir        = "info"        // 信息文件夹名
	ConfigYamlFile = "config.yaml" // 信息文件夹名

	AppListFile = "apps.json" // app列表存储文件名称
)
