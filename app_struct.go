package minimin

// App app
type App struct {
	Project
	Name string `json:"name" yaml:"name"`
	Path string `json:"path,omitempty" yaml:"path,omitempty"`
}

// AppList 每个工程目录下的appList
type AppList map[string]AppInfo

// AppInfo app 信息
type AppInfo struct {
	Name string `json:"name"`
}

// AppAdminConfig app管理的设置
type AppAdminConfig struct {
	MaxStorageDays  int `json:"max_storage_days,omitempty"`
	MaxStorageTimes int `json:"max_storage_times,omitempty"`
}
