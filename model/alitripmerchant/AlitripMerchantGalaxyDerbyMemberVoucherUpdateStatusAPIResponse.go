package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIResponse 前端订单支付成功回调-修改订单状态 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.update.status
//
// 前端订单支付成功回调-修改订单状态
type AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIResponseModel
}

// AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIResponseModel is 前端订单支付成功回调-修改订单状态 成功返回结果
type AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_update_status_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusResponse `json:"result,omitempty" xml:"result,omitempty"`
}
