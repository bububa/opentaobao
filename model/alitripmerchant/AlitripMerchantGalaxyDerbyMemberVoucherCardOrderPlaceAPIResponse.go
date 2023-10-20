package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyderbymembervouchercardorderplaceAPIResponse 德比付费会员卡下单 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.card.order.place
//
// 德比付费会员卡下单
type AlitripmerchantgalaxyderbymembervouchercardorderplaceAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxyderbymembervouchercardorderplaceAPIResponseModel
}

// AlitripmerchantgalaxyderbymembervouchercardorderplaceAPIResponseModel is 德比付费会员卡下单 成功返回结果
type AlitripmerchantgalaxyderbymembervouchercardorderplaceAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_card_order_place_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1
	Result *AlitripmerchantgalaxyderbymembervouchercardorderplaceResponse `json:"result,omitempty" xml:"result,omitempty"`
}
