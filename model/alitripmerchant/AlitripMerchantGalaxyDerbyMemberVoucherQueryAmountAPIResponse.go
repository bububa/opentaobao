package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIResponse 查询用户拥有的臻享卡数量 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.query.amount
//
// 查询用户拥有的臻享卡数量
type AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIResponseModel is 查询用户拥有的臻享卡数量 成功返回结果
type AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_query_amount_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIResponse() *AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIResponse 将 AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIResponse.Put(v)
}
