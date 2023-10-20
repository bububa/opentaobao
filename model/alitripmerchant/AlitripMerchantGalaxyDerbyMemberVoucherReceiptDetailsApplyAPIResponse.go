package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIResponse v5.0付费会员卡开发订单开票详情申请 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.receipt.details.apply
//
// v5.0付费会员卡开发订单开票详情申请
type AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIResponseModel is v5.0付费会员卡开发订单开票详情申请 成功返回结果
type AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_receipt_details_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIResponse() *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIResponse 将 AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIResponse.Put(v)
}
