package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyOrderValidateAPIResponse 星河-订单试单接口 API返回值
// alitrip.merchant.galaxy.order.validate
//
// 根据用户选择酒店房型、入住人数、预订时间参数，获取是否可预订及价格变化信息
type AlitripMerchantGalaxyOrderValidateAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyOrderValidateAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyOrderValidateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyOrderValidateAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyOrderValidateAPIResponseModel is 星河-订单试单接口 成功返回结果
type AlitripMerchantGalaxyOrderValidateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_order_validate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyOrderValidateResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyOrderValidateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyOrderValidateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyOrderValidateAPIResponse)
	},
}

// GetAlitripMerchantGalaxyOrderValidateAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyOrderValidateAPIResponse
func GetAlitripMerchantGalaxyOrderValidateAPIResponse() *AlitripMerchantGalaxyOrderValidateAPIResponse {
	return poolAlitripMerchantGalaxyOrderValidateAPIResponse.Get().(*AlitripMerchantGalaxyOrderValidateAPIResponse)
}

// ReleaseAlitripMerchantGalaxyOrderValidateAPIResponse 将 AlitripMerchantGalaxyOrderValidateAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyOrderValidateAPIResponse(v *AlitripMerchantGalaxyOrderValidateAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyOrderValidateAPIResponse.Put(v)
}
