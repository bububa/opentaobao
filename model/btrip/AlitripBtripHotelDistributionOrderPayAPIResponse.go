package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripHotelDistributionOrderPayAPIResponse 商旅酒店分销订单支付 API返回值
// alitrip.btrip.hotel.distribution.order.pay
//
// 商旅酒店分销订单支付
type AlitripBtripHotelDistributionOrderPayAPIResponse struct {
	model.CommonResponse
	AlitripBtripHotelDistributionOrderPayAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripHotelDistributionOrderPayAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripHotelDistributionOrderPayAPIResponseModel).Reset()
}

// AlitripBtripHotelDistributionOrderPayAPIResponseModel is 商旅酒店分销订单支付 成功返回结果
type AlitripBtripHotelDistributionOrderPayAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_hotel_distribution_order_pay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 支付结果返回信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 支付结果返回码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 是否支付成功
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripHotelDistributionOrderPayAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.ResultCode = 0
	m.Module = false
}

var poolAlitripBtripHotelDistributionOrderPayAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripHotelDistributionOrderPayAPIResponse)
	},
}

// GetAlitripBtripHotelDistributionOrderPayAPIResponse 从 sync.Pool 获取 AlitripBtripHotelDistributionOrderPayAPIResponse
func GetAlitripBtripHotelDistributionOrderPayAPIResponse() *AlitripBtripHotelDistributionOrderPayAPIResponse {
	return poolAlitripBtripHotelDistributionOrderPayAPIResponse.Get().(*AlitripBtripHotelDistributionOrderPayAPIResponse)
}

// ReleaseAlitripBtripHotelDistributionOrderPayAPIResponse 将 AlitripBtripHotelDistributionOrderPayAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripHotelDistributionOrderPayAPIResponse(v *AlitripBtripHotelDistributionOrderPayAPIResponse) {
	v.Reset()
	poolAlitripBtripHotelDistributionOrderPayAPIResponse.Put(v)
}
