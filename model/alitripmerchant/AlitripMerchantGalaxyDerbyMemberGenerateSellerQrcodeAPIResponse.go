package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIResponse 生成臻享卡德比分销二维码 API返回值
// alitrip.merchant.galaxy.derby.member.generate.seller.qrcode
//
// 生成臻享卡德比分销二维码
type AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIResponseModel
}

// AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIResponseModel is 生成臻享卡德比分销二维码 成功返回结果
type AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_generate_seller_qrcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeResponse `json:"result,omitempty" xml:"result,omitempty"`
}
