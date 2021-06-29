package tmallgeniescp

// SalesForecastRequest 
type SalesForecastRequest struct {
    // 入参List
    SalesForecastParamDTOList   []SalesForecastParamDTO `json:"sales_forecast_param_d_t_o_list,omitempty" xml:"sales_forecast_param_d_t_o_list>sales_forecast_param_dto,omitempty"`
    // 扩展参数
    RequestExtendJson   string `json:"request_extend_json,omitempty" xml:"request_extend_json,omitempty"`
}
