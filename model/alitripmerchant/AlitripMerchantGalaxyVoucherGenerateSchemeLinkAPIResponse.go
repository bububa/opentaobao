package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyVoucherGenerateSchemeLinkAPIResponse 生成短信链接 API返回值
// alitrip.merchant.galaxy.voucher.generate.scheme.link
//
// 生成微信跳转链接scheme_link
type AlitripMerchantGalaxyVoucherGenerateSchemeLinkAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyVoucherGenerateSchemeLinkAPIResponseModel
}

// AlitripMerchantGalaxyVoucherGenerateSchemeLinkAPIResponseModel is 生成短信链接 成功返回结果
type AlitripMerchantGalaxyVoucherGenerateSchemeLinkAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_voucher_generate_scheme_link_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlitripMerchantGalaxyVoucherGenerateSchemeLinkResponse `json:"result,omitempty" xml:"result,omitempty"`
}
