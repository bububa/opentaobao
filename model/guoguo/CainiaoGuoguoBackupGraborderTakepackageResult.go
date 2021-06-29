package guoguo

// CainiaoGuoguoBackupGraborderTakepackageResult 
type CainiaoGuoguoBackupGraborderTakepackageResult struct {
    // 封装返回的数据
    Data   *BackupOrderDO `json:"data,omitempty" xml:"data,omitempty"`
    // 调用状态
    StatusCode   string `json:"status_code,omitempty" xml:"status_code,omitempty"`
    // 调用状态码
    StatusMessage   string `json:"status_message,omitempty" xml:"status_message,omitempty"`
    // 接口调用正常
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
