package tmallgeniescp

// SupplierForecastRawRequest 
type SupplierForecastRawRequest struct {

    // ip地址
    
    Ip   string `json:"ip,omitempty" xml:"ip,omitempty"`
    

    // appkey
    
    Appkey   string `json:"appkey,omitempty" xml:"appkey,omitempty"`
    

    // 数据对象列表
    
    SupplierForecastRawParamDTOList   []SupplierForecastRawParamDto `json:"supplier_forecast_raw_param_d_t_o_list,omitempty" xml:"supplier_forecast_raw_param_d_t_o_list,omitempty"`
    

    // 扩展参数
    
    RequestExtendJson   string `json:"request_extend_json,omitempty" xml:"request_extend_json,omitempty"`
    

}
