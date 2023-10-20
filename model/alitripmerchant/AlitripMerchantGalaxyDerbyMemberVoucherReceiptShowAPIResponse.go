package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyderbymembervoucherreceiptshowAPIResponse v5.0付费会员卡开发发票图片展示 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.receipt.show
//
// v5.0付费会员卡开发发票图片展示
type AlitripmerchantgalaxyderbymembervoucherreceiptshowAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxyderbymembervoucherreceiptshowAPIResponseModel
}

// AlitripmerchantgalaxyderbymembervoucherreceiptshowAPIResponseModel is v5.0付费会员卡开发发票图片展示 成功返回结果
type AlitripmerchantgalaxyderbymembervoucherreceiptshowAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_receipt_show_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlitripmerchantgalaxyderbymembervoucherreceiptshowResponse `json:"result,omitempty" xml:"result,omitempty"`
}
