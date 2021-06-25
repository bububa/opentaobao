package tmallservice

// AlibabaSscSupplyplatformServicestoreSaveResult 
type AlibabaSscSupplyplatformServicestoreSaveResult struct {

    // 对外展示的错误信息
    DisplayMsg   string `json:"display_msg,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

    // 返回对象
    Value   *ServiceStoreCreateDto `json:"value,omitempty"`

    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty"`

}