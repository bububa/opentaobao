package waybill

// CloudPrintBaseResult 
type CloudPrintBaseResult struct {

    // data
    Data   string `json:"data,omitempty"`

    // errorCode
    ErrorCode   string `json:"error_code,omitempty"`

    // errorMessage
    ErrorMessage   string `json:"error_message,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

    // 数据
    Datas   []CustomAreaResult `json:"datas,omitempty"`

    // data
    ResourceList   []IsvResourceDo `json:"resource_list,omitempty"`

}
