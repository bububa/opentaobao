package axindata

// TaobaoalitriptravelfscrouteapiproductaddTopResult 结构体
type TaobaoalitriptravelfscrouteapiproductaddTopResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 业务数据
	FscRouteProductAddResponse *FscRouteProductAddResponse `json:"fsc_route_product_add_response,omitempty" xml:"fsc_route_product_add_response,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
