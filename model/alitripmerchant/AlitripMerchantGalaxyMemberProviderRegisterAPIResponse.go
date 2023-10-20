package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyMemberProviderRegisterAPIResponse 对外提供会员注册服务 API返回值
// alitrip.merchant.galaxy.member.provider.register
//
// 星河产品=对外提供注册雅高会员服务
type AlitripMerchantGalaxyMemberProviderRegisterAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyMemberProviderRegisterAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyMemberProviderRegisterAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyMemberProviderRegisterAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyMemberProviderRegisterAPIResponseModel is 对外提供会员注册服务 成功返回结果
type AlitripMerchantGalaxyMemberProviderRegisterAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_member_provider_register_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyMemberProviderRegisterResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyMemberProviderRegisterAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyMemberProviderRegisterAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyMemberProviderRegisterAPIResponse)
	},
}

// GetAlitripMerchantGalaxyMemberProviderRegisterAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyMemberProviderRegisterAPIResponse
func GetAlitripMerchantGalaxyMemberProviderRegisterAPIResponse() *AlitripMerchantGalaxyMemberProviderRegisterAPIResponse {
	return poolAlitripMerchantGalaxyMemberProviderRegisterAPIResponse.Get().(*AlitripMerchantGalaxyMemberProviderRegisterAPIResponse)
}

// ReleaseAlitripMerchantGalaxyMemberProviderRegisterAPIResponse 将 AlitripMerchantGalaxyMemberProviderRegisterAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyMemberProviderRegisterAPIResponse(v *AlitripMerchantGalaxyMemberProviderRegisterAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyMemberProviderRegisterAPIResponse.Put(v)
}
