package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardActiveAPIResponse 权益卡订单激活 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.card.active
//
// 权益卡订单激活
type AlitripMerchantGalaxyDerbyMemberVoucherCardActiveAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyDerbyMemberVoucherCardActiveAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherCardActiveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyDerbyMemberVoucherCardActiveAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyDerbyMemberVoucherCardActiveAPIResponseModel is 权益卡订单激活 成功返回结果
type AlitripMerchantGalaxyDerbyMemberVoucherCardActiveAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_card_active_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlitripMerchantGalaxyDerbyMemberVoucherCardActiveResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherCardActiveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherCardActiveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherCardActiveAPIResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherCardActiveAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyMemberVoucherCardActiveAPIResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherCardActiveAPIResponse() *AlitripMerchantGalaxyDerbyMemberVoucherCardActiveAPIResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherCardActiveAPIResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherCardActiveAPIResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardActiveAPIResponse 将 AlitripMerchantGalaxyDerbyMemberVoucherCardActiveAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardActiveAPIResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherCardActiveAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyMemberVoucherCardActiveAPIResponse.Put(v)
}
