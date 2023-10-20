package alitripmerchant

import (
	"encoding/xml"

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

// AlitripMerchantGalaxyCommonGetEnumsbynameAPIResponseModel is 小程序公共枚举类获取接口 成功返回结果
type AlitripMerchantGalaxyCommonGetEnumsbynameAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_common_get_enumsbyname_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 携带优惠券试单结果
	Result *AlitripMerchantGalaxyCommonGetEnumsbynameResponse `json:"result,omitempty" xml:"result,omitempty"`
}
