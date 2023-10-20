package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyderbymembervouchercardactiveAPIResponse 权益卡订单激活 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.card.active
//
// 权益卡订单激活
type AlitripmerchantgalaxyderbymembervouchercardactiveAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxyderbymembervouchercardactiveAPIResponseModel
}

// AlitripmerchantgalaxyderbymembervouchercardactiveAPIResponseModel is 权益卡订单激活 成功返回结果
type AlitripmerchantgalaxyderbymembervouchercardactiveAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_card_active_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlitripmerchantgalaxyderbymembervouchercardactiveResponse `json:"result,omitempty" xml:"result,omitempty"`
}
