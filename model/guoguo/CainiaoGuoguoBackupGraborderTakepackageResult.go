package guoguo

// CainiaoGuoguoBackupGraborderTakepackageResult 
/* model for simplify = false
type CainiaoGuoguoBackupGraborderTakepackageResult struct {

    // 封装返回的数据
    
    Data  *struct {
        BackupOrderDo  *BackupOrderDo `json:"backup_order_do,omitempty"`
    } `json:"data,omitempty"`
    

    // 调用状态
    
    StatusCode   string `json:"status_code,omitempty"`
    

    // 调用状态码
    
    StatusMessage   string `json:"status_message,omitempty"`
    

    // 接口调用正常
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// CainiaoGuoguoBackupGraborderTakepackageResult 
type CainiaoGuoguoBackupGraborderTakepackageResult struct {

    // 封装返回的数据
    Data   *BackupOrderDo `json:"data,omitempty"`

    // 调用状态
    StatusCode   string `json:"status_code,omitempty"`

    // 调用状态码
    StatusMessage   string `json:"status_message,omitempty"`

    // 接口调用正常
    Success   bool `json:"success,omitempty"`

}
