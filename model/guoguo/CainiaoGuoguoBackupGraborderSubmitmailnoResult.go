package guoguo

// CainiaoGuoguoBackupGraborderSubmitmailnoResult 
type CainiaoGuoguoBackupGraborderSubmitmailnoResult struct {
    // 数据对象
    Data   *BackupOrderDO `json:"data,omitempty" xml:"data,omitempty"`
    // 1
    StatusCode   string `json:"status_code,omitempty" xml:"status_code,omitempty"`
    // 返回的状态描述
    StatusMessage   string `json:"status_message,omitempty" xml:"status_message,omitempty"`
    // 成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
