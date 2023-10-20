package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyderbymembervoucherreceiptautoAPIResponse 德比付费会员卡开票自匹配 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.receipt.auto
//
// 德比付费会员卡开票自匹配
type AlitripmerchantgalaxyderbymembervoucherreceiptautoAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxyderbymembervoucherreceiptautoAPIResponseModel
}

// AlitripmerchantgalaxyderbymembervoucherreceiptautoAPIResponseModel is 德比付费会员卡开票自匹配 成功返回结果
type AlitripmerchantgalaxyderbymembervoucherreceiptautoAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_receipt_auto_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlitripmerchantgalaxyderbymembervoucherreceiptautoResponse `json:"result,omitempty" xml:"result,omitempty"`
}
