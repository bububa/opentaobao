package btrip

// HisvResultSupport 
type HisvResultSupport struct {
    // 请求是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 出参
    ModuleList   []OpenApiZzdFlightOrderRs `json:"module_list,omitempty" xml:"module_list>open_api_zzd_flight_order_rs,omitempty"`
    // 错误码
    ResultCode   int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // 错误信息
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
    // 出参
    TradeList   []OpenApiZzdHotelOrderRs `json:"trade_list,omitempty" xml:"trade_list>open_api_zzd_hotel_order_rs,omitempty"`
}
