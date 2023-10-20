package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoAPIResponse v5.0付费会员卡开票抬头自匹配 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.receipt.details.auto
//
// v5.0付费会员卡开票抬头自匹配
type AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoAPIResponseModel is v5.0付费会员卡开票抬头自匹配 成功返回结果
type AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_receipt_details_auto_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 抬头自匹配接口返回结果
	Result *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoAPIResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoAPIResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoAPIResponse() *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoAPIResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoAPIResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoAPIResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoAPIResponse 将 AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoAPIResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoAPIResponse.Put(v)
}
