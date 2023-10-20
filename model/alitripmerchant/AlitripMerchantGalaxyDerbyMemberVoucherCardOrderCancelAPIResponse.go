package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyderbymembervouchercardordercancelAPIResponse 取消订单 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.card.order.cancel
//
// 德比付费会员卡订单取消
type AlitripmerchantgalaxyderbymembervouchercardordercancelAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxyderbymembervouchercardordercancelAPIResponseModel
}

// AlitripmerchantgalaxyderbymembervouchercardordercancelAPIResponseModel is 取消订单 成功返回结果
type AlitripmerchantgalaxyderbymembervouchercardordercancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_card_order_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1
	Result *AlitripmerchantgalaxyderbymembervouchercardordercancelResponse `json:"result,omitempty" xml:"result,omitempty"`
}
