package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripHotelDistributionOrderCancelAPIResponse 商旅酒店API分销取消订单 API返回值
// alitrip.btrip.hotel.distribution.order.cancel
//
// 商旅酒店API分销取消订单
type AlitripBtripHotelDistributionOrderCancelAPIResponse struct {
	model.CommonResponse
	AlitripBtripHotelDistributionOrderCancelAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripHotelDistributionOrderCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripHotelDistributionOrderCancelAPIResponseModel).Reset()
}

// AlitripBtripHotelDistributionOrderCancelAPIResponseModel is 商旅酒店API分销取消订单 成功返回结果
type AlitripBtripHotelDistributionOrderCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_hotel_distribution_order_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 取消订单返回结果
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripHotelDistributionOrderCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripHotelDistributionOrderCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripHotelDistributionOrderCancelAPIResponse)
	},
}

// GetAlitripBtripHotelDistributionOrderCancelAPIResponse 从 sync.Pool 获取 AlitripBtripHotelDistributionOrderCancelAPIResponse
func GetAlitripBtripHotelDistributionOrderCancelAPIResponse() *AlitripBtripHotelDistributionOrderCancelAPIResponse {
	return poolAlitripBtripHotelDistributionOrderCancelAPIResponse.Get().(*AlitripBtripHotelDistributionOrderCancelAPIResponse)
}

// ReleaseAlitripBtripHotelDistributionOrderCancelAPIResponse 将 AlitripBtripHotelDistributionOrderCancelAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripHotelDistributionOrderCancelAPIResponse(v *AlitripBtripHotelDistributionOrderCancelAPIResponse) {
	v.Reset()
	poolAlitripBtripHotelDistributionOrderCancelAPIResponse.Put(v)
}
