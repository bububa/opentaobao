package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIResponse Aps退券通知接口 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.card.aps.refund
//
// Aps退券通知接口
type AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIResponseModel is Aps退券通知接口 成功返回结果
type AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_card_aps_refund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应
	Result *AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIResponse() *AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIResponse 将 AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIResponse.Put(v)
}
