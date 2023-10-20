package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyderbymembervouchercardchangecallbackAPIResponse v5.0德比付费会员卡通知 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.card.change.callback
//
// v5.0德比付费会员卡通知
type AlitripmerchantgalaxyderbymembervouchercardchangecallbackAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxyderbymembervouchercardchangecallbackAPIResponseModel
}

// AlitripmerchantgalaxyderbymembervouchercardchangecallbackAPIResponseModel is v5.0德比付费会员卡通知 成功返回结果
type AlitripmerchantgalaxyderbymembervouchercardchangecallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_card_change_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1
	Result *AlitripmerchantgalaxyderbymembervouchercardchangecallbackResponse `json:"result,omitempty" xml:"result,omitempty"`
}
