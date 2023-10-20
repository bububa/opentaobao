package guoguo

// CainiaoGuoguoBackupGraborderTakepackageResult 结构体
type CainiaoGuoguoBackupGraborderTakepackageResult struct {
	// 调用状态
	StatusCode string `json:"status_code,omitempty" xml:"status_code,omitempty"`
	// 调用状态码
	StatusMessage string `json:"status_message,omitempty" xml:"status_message,omitempty"`
	// 封装返回的数据
	Data *BackupOrderDo `json:"data,omitempty" xml:"data,omitempty"`
	// 接口调用正常
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
