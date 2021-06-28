package guoguo

// CainiaoGuoguoBackupGraborderSubmitmailnoResult 
/* model for simplify = false
type CainiaoGuoguoBackupGraborderSubmitmailnoResult struct {

    // 数据对象
    
    Data  *struct {
        BackupOrderDo  *BackupOrderDo `json:"backup_order_do,omitempty"`
    } `json:"data,omitempty"`
    

    // 1
    
    StatusCode   string `json:"status_code,omitempty"`
    

    // 返回的状态描述
    
    StatusMessage   string `json:"status_message,omitempty"`
    

    // 成功
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// CainiaoGuoguoBackupGraborderSubmitmailnoResult 
type CainiaoGuoguoBackupGraborderSubmitmailnoResult struct {

    // 数据对象
    Data   *BackupOrderDo `json:"data,omitempty"`

    // 1
    StatusCode   string `json:"status_code,omitempty"`

    // 返回的状态描述
    StatusMessage   string `json:"status_message,omitempty"`

    // 成功
    Success   bool `json:"success,omitempty"`

}
