package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIResponse 德比付费会员卡查询 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.card.query
//
// 德比付费会员卡查询
type AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIResponseModel is 德比付费会员卡查询 成功返回结果
type AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_card_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1
	Result *AlitripMerchantGalaxyDerbyMemberVoucherCardQueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIResponse() *AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIResponse 将 AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIResponse.Put(v)
}
