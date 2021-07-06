package guoguo

// CainiaoGuoguoBackupGraborderSubmitmailnoResult 结构体
type CainiaoGuoguoBackupGraborderSubmitmailnoResult struct {
	// 1
	StatusCode string `json:"status_code,omitempty" xml:"status_code,omitempty"`
	// 返回的状态描述
	StatusMessage string `json:"status_message,omitempty" xml:"status_message,omitempty"`
	// 数据对象
	Data *BackupOrderDo `json:"data,omitempty" xml:"data,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
