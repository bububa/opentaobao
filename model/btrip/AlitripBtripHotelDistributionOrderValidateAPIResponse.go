package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripHotelDistributionOrderValidateAPIResponse 商旅酒店API分销下单前校验 API返回值
// alitrip.btrip.hotel.distribution.order.validate
//
// 商旅酒店API分销下单前校验
type AlitripBtripHotelDistributionOrderValidateAPIResponse struct {
	model.CommonResponse
	AlitripBtripHotelDistributionOrderValidateAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripHotelDistributionOrderValidateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripHotelDistributionOrderValidateAPIResponseModel).Reset()
}

// AlitripBtripHotelDistributionOrderValidateAPIResponseModel is 商旅酒店API分销下单前校验 成功返回结果
type AlitripBtripHotelDistributionOrderValidateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_hotel_distribution_order_validate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 下单前校验结果
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripHotelDistributionOrderValidateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripHotelDistributionOrderValidateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripHotelDistributionOrderValidateAPIResponse)
	},
}

// GetAlitripBtripHotelDistributionOrderValidateAPIResponse 从 sync.Pool 获取 AlitripBtripHotelDistributionOrderValidateAPIResponse
func GetAlitripBtripHotelDistributionOrderValidateAPIResponse() *AlitripBtripHotelDistributionOrderValidateAPIResponse {
	return poolAlitripBtripHotelDistributionOrderValidateAPIResponse.Get().(*AlitripBtripHotelDistributionOrderValidateAPIResponse)
}

// ReleaseAlitripBtripHotelDistributionOrderValidateAPIResponse 将 AlitripBtripHotelDistributionOrderValidateAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripHotelDistributionOrderValidateAPIResponse(v *AlitripBtripHotelDistributionOrderValidateAPIResponse) {
	v.Reset()
	poolAlitripBtripHotelDistributionOrderValidateAPIResponse.Put(v)
}
