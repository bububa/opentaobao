package tmallgeniescp

// SupplierForecastRequest 
type SupplierForecastRequest struct {

    // 请求参数
    
    SupplierForecastParamDTOList   []SupplierForecastParamDto `json:"supplier_forecast_param_d_t_o_list,omitempty" xml:"supplier_forecast_param_d_t_o_list,omitempty"`
    

    // 扩展参数
    
    RequestExtendJson   string `json:"request_extend_json,omitempty" xml:"request_extend_json,omitempty"`
    

}
