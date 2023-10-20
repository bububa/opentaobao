package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyderbymembervoucherreceiptdetailsqueryAPIResponse v5.0付费会员卡开发订单开票信息查询 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.receipt.details.query
//
// v5.0付费会员卡开发订单开票信息查询
type AlitripmerchantgalaxyderbymembervoucherreceiptdetailsqueryAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxyderbymembervoucherreceiptdetailsqueryAPIResponseModel
}

// AlitripmerchantgalaxyderbymembervoucherreceiptdetailsqueryAPIResponseModel is v5.0付费会员卡开发订单开票信息查询 成功返回结果
type AlitripmerchantgalaxyderbymembervoucherreceiptdetailsqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_receipt_details_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *AlitripmerchantgalaxyderbymembervoucherreceiptdetailsqueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}
