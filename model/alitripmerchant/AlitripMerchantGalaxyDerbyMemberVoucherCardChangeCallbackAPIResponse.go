package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIResponse v5.0德比付费会员卡通知 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.card.change.callback
//
// v5.0德比付费会员卡通知
type AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIResponseModel is v5.0德比付费会员卡通知 成功返回结果
type AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_card_change_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1
	Result *AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIResponse() *AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIResponse 将 AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIResponse.Put(v)
}
