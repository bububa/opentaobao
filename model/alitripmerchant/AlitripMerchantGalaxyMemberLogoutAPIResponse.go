package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyMemberLogoutAPIResponse 星河-用户登出 API返回值
// alitrip.merchant.galaxy.member.logout
//
// 星河=微信小程序用户登出
type AlitripMerchantGalaxyMemberLogoutAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyMemberLogoutAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyMemberLogoutAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyMemberLogoutAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyMemberLogoutAPIResponseModel is 星河-用户登出 成功返回结果
type AlitripMerchantGalaxyMemberLogoutAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_member_logout_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyMemberLogoutResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyMemberLogoutAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyMemberLogoutAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyMemberLogoutAPIResponse)
	},
}

// GetAlitripMerchantGalaxyMemberLogoutAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyMemberLogoutAPIResponse
func GetAlitripMerchantGalaxyMemberLogoutAPIResponse() *AlitripMerchantGalaxyMemberLogoutAPIResponse {
	return poolAlitripMerchantGalaxyMemberLogoutAPIResponse.Get().(*AlitripMerchantGalaxyMemberLogoutAPIResponse)
}

// ReleaseAlitripMerchantGalaxyMemberLogoutAPIResponse 将 AlitripMerchantGalaxyMemberLogoutAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyMemberLogoutAPIResponse(v *AlitripMerchantGalaxyMemberLogoutAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyMemberLogoutAPIResponse.Put(v)
}
