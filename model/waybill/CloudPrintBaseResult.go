package waybill

// CloudPrintBaseResult 
type CloudPrintBaseResult struct {
    // data
    Data   string `json:"data,omitempty" xml:"data,omitempty"`
    // errorCode
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // errorMessage
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    // success
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 数据
    Datas   []CustomAreaResult `json:"datas,omitempty" xml:"datas>custom_area_result,omitempty"`
    // data
    ResourceList   []IsvResourceDO `json:"resource_list,omitempty" xml:"resource_list>isv_resource_do,omitempty"`
}
