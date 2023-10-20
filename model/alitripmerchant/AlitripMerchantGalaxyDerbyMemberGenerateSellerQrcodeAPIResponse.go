package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyderbymembergeneratesellerqrcodeAPIResponse 生成臻享卡德比分销二维码 API返回值
// alitrip.merchant.galaxy.derby.member.generate.seller.qrcode
//
// 生成臻享卡德比分销二维码
type AlitripmerchantgalaxyderbymembergeneratesellerqrcodeAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxyderbymembergeneratesellerqrcodeAPIResponseModel
}

// AlitripmerchantgalaxyderbymembergeneratesellerqrcodeAPIResponseModel is 生成臻享卡德比分销二维码 成功返回结果
type AlitripmerchantgalaxyderbymembergeneratesellerqrcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_generate_seller_qrcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlitripmerchantgalaxyderbymembergeneratesellerqrcodeResponse `json:"result,omitempty" xml:"result,omitempty"`
}
