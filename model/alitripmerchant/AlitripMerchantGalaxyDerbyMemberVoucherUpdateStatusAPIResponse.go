package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIResponse 前端订单支付成功回调-修改订单状态 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.update.status
//
// 前端订单支付成功回调-修改订单状态
type AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIResponseModel is 前端订单支付成功回调-修改订单状态 成功返回结果
type AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_update_status_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIResponse() *AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIResponse 将 AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIResponse.Put(v)
}
