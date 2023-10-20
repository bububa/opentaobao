package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyvouchergenerateschemelinkAPIResponse 生成短信链接 API返回值
// alitrip.merchant.galaxy.voucher.generate.scheme.link
//
// 生成微信跳转链接scheme_link
type AlitripmerchantgalaxyvouchergenerateschemelinkAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxyvouchergenerateschemelinkAPIResponseModel
}

// AlitripmerchantgalaxyvouchergenerateschemelinkAPIResponseModel is 生成短信链接 成功返回结果
type AlitripmerchantgalaxyvouchergenerateschemelinkAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_voucher_generate_scheme_link_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlitripmerchantgalaxyvouchergenerateschemelinkResponse `json:"result,omitempty" xml:"result,omitempty"`
}
