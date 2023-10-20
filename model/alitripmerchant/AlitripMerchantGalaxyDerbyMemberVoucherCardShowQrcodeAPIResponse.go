package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIResponse 会员权益卡身份识别二维码图片 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.card.show.qrcode
//
// 会员权益卡身份识别二维码图片
type AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIResponseModel
}

// AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIResponseModel is 会员权益卡身份识别二维码图片 成功返回结果
type AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_card_show_qrcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// response
	Result *AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeResponse `json:"result,omitempty" xml:"result,omitempty"`
}
