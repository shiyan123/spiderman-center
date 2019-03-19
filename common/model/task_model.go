package model

type TaskInfo struct {
	TaskId   string
	TaskName string
	Config   *TaskConfig
}

type TaskConfig struct {
	TaskType        int    //任务类型
	CreatedAt       int64  //创建时间
	LastBeginAt     int64  //上一次开始时间
	LastEndAt       int64  //上一次结束时间
	ProgramUpdateAt int64  //任务程序更新时间
	ProgramName     string //运行程序
	CronStr         string //cron 表达式
	Path            string //运行目录
	LogPath         string //日志目录
	Status          int    //运行状态
	Retry           int    //重试次数
}
