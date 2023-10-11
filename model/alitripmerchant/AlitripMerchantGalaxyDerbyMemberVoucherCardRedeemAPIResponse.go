package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIResponse 根据兑换码兑换臻享卡接口 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.card.redeem
//
// 根据兑换码兑换臻享卡接口
type AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIResponseModel
}

// AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIResponseModel is 根据兑换码兑换臻享卡接口 成功返回结果
type AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_card_redeem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// response
	Result *AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemResponse `json:"result,omitempty" xml:"result,omitempty"`
}
