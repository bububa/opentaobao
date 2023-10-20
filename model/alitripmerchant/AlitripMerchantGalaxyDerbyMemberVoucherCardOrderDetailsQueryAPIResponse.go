package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIResponse 订单详情查询接口 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.card.order.details.query
//
// 订单详情查询接口
type AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIResponseModel is 订单详情查询接口 成功返回结果
type AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_card_order_details_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIResponse() *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIResponse 将 AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIResponse.Put(v)
}
