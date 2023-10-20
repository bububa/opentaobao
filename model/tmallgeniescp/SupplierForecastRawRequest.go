package tmallgeniescp

import (
	"sync"
)

// SupplierForecastRawRequest 结构体
type SupplierForecastRawRequest struct {
	// 数据对象列表
	SupplierForecastRawParamDTOList []SupplierForecastRawParamDto `json:"supplier_forecast_raw_param_d_t_o_list,omitempty" xml:"supplier_forecast_raw_param_d_t_o_list>supplier_forecast_raw_param_dto,omitempty"`
	// ip地址
	Ip string `json:"ip,omitempty" xml:"ip,omitempty"`
	// appkey
	Appkey string `json:"appkey,omitempty" xml:"appkey,omitempty"`
	// 扩展参数
	RequestExtendJson string `json:"request_extend_json,omitempty" xml:"request_extend_json,omitempty"`
}

var poolSupplierForecastRawRequest = sync.Pool{
	New: func() any {
		return new(SupplierForecastRawRequest)
	},
}

// GetSupplierForecastRawRequest() 从对象池中获取SupplierForecastRawRequest
func GetSupplierForecastRawRequest() *SupplierForecastRawRequest {
	return poolSupplierForecastRawRequest.Get().(*SupplierForecastRawRequest)
}

// ReleaseSupplierForecastRawRequest 释放SupplierForecastRawRequest
func ReleaseSupplierForecastRawRequest(v *SupplierForecastRawRequest) {
	v.SupplierForecastRawParamDTOList = v.SupplierForecastRawParamDTOList[:0]
	v.Ip = ""
	v.Appkey = ""
	v.RequestExtendJson = ""
	poolSupplierForecastRawRequest.Put(v)
}
