package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIResponse 查询会员兑换臻享卡历史记录 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.card.redeem.history
//
// 查询会员兑换臻享卡历史记录
type AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIResponseModel is 查询会员兑换臻享卡历史记录 成功返回结果
type AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_card_redeem_history_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// response
	Result *AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIResponse() *AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIResponse 将 AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIResponse.Put(v)
}
