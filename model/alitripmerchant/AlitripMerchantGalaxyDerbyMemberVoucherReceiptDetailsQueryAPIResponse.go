package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIResponse v5.0付费会员卡开发订单开票信息查询 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.receipt.details.query
//
// v5.0付费会员卡开发订单开票信息查询
type AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIResponseModel is v5.0付费会员卡开发订单开票信息查询 成功返回结果
type AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_receipt_details_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIResponse() *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIResponse 将 AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIResponse.Put(v)
}
