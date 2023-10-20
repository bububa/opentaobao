package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIResponse 查询权益卡订单列表 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.card.orders.query
//
// 查询权益卡订单列表
type AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIResponseModel is 查询权益卡订单列表 成功返回结果
type AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_card_orders_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIResponse() *AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIResponse 将 AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIResponse.Put(v)
}
