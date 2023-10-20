package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripHotelDistributionOrderCreateAPIResponse 商旅酒店分销-创建订单 API返回值
// alitrip.btrip.hotel.distribution.order.create
//
// 商旅酒店分销-创建订单
type AlitripBtripHotelDistributionOrderCreateAPIResponse struct {
	model.CommonResponse
	AlitripBtripHotelDistributionOrderCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripHotelDistributionOrderCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripHotelDistributionOrderCreateAPIResponseModel).Reset()
}

// AlitripBtripHotelDistributionOrderCreateAPIResponseModel is 商旅酒店分销-创建订单 成功返回结果
type AlitripBtripHotelDistributionOrderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_hotel_distribution_order_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创单返回结果
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripHotelDistributionOrderCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripHotelDistributionOrderCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripHotelDistributionOrderCreateAPIResponse)
	},
}

// GetAlitripBtripHotelDistributionOrderCreateAPIResponse 从 sync.Pool 获取 AlitripBtripHotelDistributionOrderCreateAPIResponse
func GetAlitripBtripHotelDistributionOrderCreateAPIResponse() *AlitripBtripHotelDistributionOrderCreateAPIResponse {
	return poolAlitripBtripHotelDistributionOrderCreateAPIResponse.Get().(*AlitripBtripHotelDistributionOrderCreateAPIResponse)
}

// ReleaseAlitripBtripHotelDistributionOrderCreateAPIResponse 将 AlitripBtripHotelDistributionOrderCreateAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripHotelDistributionOrderCreateAPIResponse(v *AlitripBtripHotelDistributionOrderCreateAPIResponse) {
	v.Reset()
	poolAlitripBtripHotelDistributionOrderCreateAPIResponse.Put(v)
}
