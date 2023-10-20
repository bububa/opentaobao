package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIResponse v5.0付费会员卡开发发票图片展示 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.receipt.show
//
// v5.0付费会员卡开发发票图片展示
type AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIResponseModel is v5.0付费会员卡开发发票图片展示 成功返回结果
type AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_receipt_show_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIResponse() *AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIResponse 将 AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIResponse.Put(v)
}
