package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIResponse 取消订单 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.card.order.cancel
//
// 德比付费会员卡订单取消
type AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIResponseModel is 取消订单 成功返回结果
type AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_card_order_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1
	Result *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIResponse() *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIResponse 将 AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIResponse.Put(v)
}
