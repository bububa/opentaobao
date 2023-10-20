package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIResponse 德比付费会员卡下单 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.card.order.place
//
// 德比付费会员卡下单
type AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIResponseModel is 德比付费会员卡下单 成功返回结果
type AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_card_order_place_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1
	Result *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIResponse() *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIResponse 将 AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIResponse.Put(v)
}
