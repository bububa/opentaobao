package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyderbymembervouchercardordersqueryAPIResponse 查询权益卡订单列表 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.card.orders.query
//
// 查询权益卡订单列表
type AlitripmerchantgalaxyderbymembervouchercardordersqueryAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxyderbymembervouchercardordersqueryAPIResponseModel
}

// AlitripmerchantgalaxyderbymembervouchercardordersqueryAPIResponseModel is 查询权益卡订单列表 成功返回结果
type AlitripmerchantgalaxyderbymembervouchercardordersqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_card_orders_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlitripmerchantgalaxyderbymembervouchercardordersqueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}
