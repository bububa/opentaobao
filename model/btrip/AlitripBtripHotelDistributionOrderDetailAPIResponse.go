package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripHotelDistributionOrderDetailAPIResponse 商旅酒店API分销查询订单详情 API返回值
// alitrip.btrip.hotel.distribution.order.detail
//
// 商旅酒店API分销查询订单详情
type AlitripBtripHotelDistributionOrderDetailAPIResponse struct {
	model.CommonResponse
	AlitripBtripHotelDistributionOrderDetailAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripHotelDistributionOrderDetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripHotelDistributionOrderDetailAPIResponseModel).Reset()
}

// AlitripBtripHotelDistributionOrderDetailAPIResponseModel is 商旅酒店API分销查询订单详情 成功返回结果
type AlitripBtripHotelDistributionOrderDetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_hotel_distribution_order_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单详情接口返回结果
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripHotelDistributionOrderDetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripHotelDistributionOrderDetailAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripHotelDistributionOrderDetailAPIResponse)
	},
}

// GetAlitripBtripHotelDistributionOrderDetailAPIResponse 从 sync.Pool 获取 AlitripBtripHotelDistributionOrderDetailAPIResponse
func GetAlitripBtripHotelDistributionOrderDetailAPIResponse() *AlitripBtripHotelDistributionOrderDetailAPIResponse {
	return poolAlitripBtripHotelDistributionOrderDetailAPIResponse.Get().(*AlitripBtripHotelDistributionOrderDetailAPIResponse)
}

// ReleaseAlitripBtripHotelDistributionOrderDetailAPIResponse 将 AlitripBtripHotelDistributionOrderDetailAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripHotelDistributionOrderDetailAPIResponse(v *AlitripBtripHotelDistributionOrderDetailAPIResponse) {
	v.Reset()
	poolAlitripBtripHotelDistributionOrderDetailAPIResponse.Put(v)
}
