package tmallgeniescp

import (
	"sync"
)

// SalesForecastRequest 结构体
type SalesForecastRequest struct {
	// 入参List
	SalesForecastParamDTOList []SalesForecastParamDto `json:"sales_forecast_param_d_t_o_list,omitempty" xml:"sales_forecast_param_d_t_o_list>sales_forecast_param_dto,omitempty"`
	// 扩展参数
	RequestExtendJson string `json:"request_extend_json,omitempty" xml:"request_extend_json,omitempty"`
}

var poolSalesForecastRequest = sync.Pool{
	New: func() any {
		return new(SalesForecastRequest)
	},
}

// GetSalesForecastRequest() 从对象池中获取SalesForecastRequest
func GetSalesForecastRequest() *SalesForecastRequest {
	return poolSalesForecastRequest.Get().(*SalesForecastRequest)
}

// ReleaseSalesForecastRequest 释放SalesForecastRequest
func ReleaseSalesForecastRequest(v *SalesForecastRequest) {
	v.SalesForecastParamDTOList = v.SalesForecastParamDTOList[:0]
	v.RequestExtendJson = ""
	poolSalesForecastRequest.Put(v)
}
