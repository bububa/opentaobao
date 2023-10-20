package tmallservice

// AlibabasscsupplyplatformservicedefinitionqueryskuResult 结构体
type AlibabasscsupplyplatformservicedefinitionqueryskuResult struct {
	// 服务sku列表
	ServiceSkus []ServiceSkuDto `json:"service_skus,omitempty" xml:"service_skus>service_sku_dto,omitempty"`
	// 对外展示的错误信息
	DisplayMsg string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success string `json:"success,omitempty" xml:"success,omitempty"`
}
