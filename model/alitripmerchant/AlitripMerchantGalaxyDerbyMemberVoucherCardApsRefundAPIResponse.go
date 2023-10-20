package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyderbymembervouchercardapsrefundAPIResponse Aps退券通知接口 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.card.aps.refund
//
// Aps退券通知接口
type AlitripmerchantgalaxyderbymembervouchercardapsrefundAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxyderbymembervouchercardapsrefundAPIResponseModel
}

// AlitripmerchantgalaxyderbymembervouchercardapsrefundAPIResponseModel is Aps退券通知接口 成功返回结果
type AlitripmerchantgalaxyderbymembervouchercardapsrefundAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_card_aps_refund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应
	Result *AlitripmerchantgalaxyderbymembervouchercardapsrefundResponse `json:"result,omitempty" xml:"result,omitempty"`
}
