package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIResponse 查询会员兑换臻享卡历史记录 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.card.redeem.history
//
// 查询会员兑换臻享卡历史记录
type AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIResponseModel
}

// AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIResponseModel is 查询会员兑换臻享卡历史记录 成功返回结果
type AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_card_redeem_history_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// response
	Result *AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryResponse `json:"result,omitempty" xml:"result,omitempty"`
}
