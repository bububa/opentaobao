package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIResponse Aps退券通知接口 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.card.aps.refund
//
// Aps退券通知接口
type AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIResponseModel
}

// AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIResponseModel is Aps退券通知接口 成功返回结果
type AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_card_aps_refund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应
	Result *AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundResponse `json:"result,omitempty" xml:"result,omitempty"`
}
