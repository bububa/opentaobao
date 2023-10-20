package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyCommonGetEnumsbynameAPIResponse 小程序公共枚举类获取接口 API返回值
// alitrip.merchant.galaxy.common.get.enumsbyname
//
// 通过枚举名称获取枚举类实例内容
type AlitripMerchantGalaxyCommonGetEnumsbynameAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyCommonGetEnumsbynameAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyCommonGetEnumsbynameAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyCommonGetEnumsbynameAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyCommonGetEnumsbynameAPIResponseModel is 小程序公共枚举类获取接口 成功返回结果
type AlitripMerchantGalaxyCommonGetEnumsbynameAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_common_get_enumsbyname_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 携带优惠券试单结果
	Result *AlitripMerchantGalaxyCommonGetEnumsbynameResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyCommonGetEnumsbynameAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyCommonGetEnumsbynameAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyCommonGetEnumsbynameAPIResponse)
	},
}

// GetAlitripMerchantGalaxyCommonGetEnumsbynameAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyCommonGetEnumsbynameAPIResponse
func GetAlitripMerchantGalaxyCommonGetEnumsbynameAPIResponse() *AlitripMerchantGalaxyCommonGetEnumsbynameAPIResponse {
	return poolAlitripMerchantGalaxyCommonGetEnumsbynameAPIResponse.Get().(*AlitripMerchantGalaxyCommonGetEnumsbynameAPIResponse)
}

// ReleaseAlitripMerchantGalaxyCommonGetEnumsbynameAPIResponse 将 AlitripMerchantGalaxyCommonGetEnumsbynameAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyCommonGetEnumsbynameAPIResponse(v *AlitripMerchantGalaxyCommonGetEnumsbynameAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyCommonGetEnumsbynameAPIResponse.Put(v)
}
