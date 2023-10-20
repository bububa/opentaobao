package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyOrderFillAPIResponse 填单页接口 API返回值
// alitrip.merchant.galaxy.order.fill
//
// 进入填单页时调用此接口，返回填单所需展示基础信息
type AlitripMerchantGalaxyOrderFillAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyOrderFillAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyOrderFillAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyOrderFillAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyOrderFillAPIResponseModel is 填单页接口 成功返回结果
type AlitripMerchantGalaxyOrderFillAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_order_fill_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyOrderFillResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyOrderFillAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyOrderFillAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyOrderFillAPIResponse)
	},
}

// GetAlitripMerchantGalaxyOrderFillAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyOrderFillAPIResponse
func GetAlitripMerchantGalaxyOrderFillAPIResponse() *AlitripMerchantGalaxyOrderFillAPIResponse {
	return poolAlitripMerchantGalaxyOrderFillAPIResponse.Get().(*AlitripMerchantGalaxyOrderFillAPIResponse)
}

// ReleaseAlitripMerchantGalaxyOrderFillAPIResponse 将 AlitripMerchantGalaxyOrderFillAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyOrderFillAPIResponse(v *AlitripMerchantGalaxyOrderFillAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyOrderFillAPIResponse.Put(v)
}
