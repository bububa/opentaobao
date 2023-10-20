package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyMemberTokenAPIResponse 星河-校验token API返回值
// alitrip.merchant.galaxy.member.token
//
// 校验或者刷新token
type AlitripMerchantGalaxyMemberTokenAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyMemberTokenAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyMemberTokenAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyMemberTokenAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyMemberTokenAPIResponseModel is 星河-校验token 成功返回结果
type AlitripMerchantGalaxyMemberTokenAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_member_token_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyMemberTokenResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyMemberTokenAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyMemberTokenAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyMemberTokenAPIResponse)
	},
}

// GetAlitripMerchantGalaxyMemberTokenAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyMemberTokenAPIResponse
func GetAlitripMerchantGalaxyMemberTokenAPIResponse() *AlitripMerchantGalaxyMemberTokenAPIResponse {
	return poolAlitripMerchantGalaxyMemberTokenAPIResponse.Get().(*AlitripMerchantGalaxyMemberTokenAPIResponse)
}

// ReleaseAlitripMerchantGalaxyMemberTokenAPIResponse 将 AlitripMerchantGalaxyMemberTokenAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyMemberTokenAPIResponse(v *AlitripMerchantGalaxyMemberTokenAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyMemberTokenAPIResponse.Put(v)
}
