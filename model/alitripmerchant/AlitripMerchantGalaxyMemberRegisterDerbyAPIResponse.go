package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyMemberRegisterDerbyAPIResponse 会员注册(新版注册接口对接德比) API返回值
// alitrip.merchant.galaxy.member.register.derby
//
// 会员注册(新版注册接口对接德比)，返回手机号/邮箱/注册状态
type AlitripMerchantGalaxyMemberRegisterDerbyAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyMemberRegisterDerbyAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyMemberRegisterDerbyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyMemberRegisterDerbyAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyMemberRegisterDerbyAPIResponseModel is 会员注册(新版注册接口对接德比) 成功返回结果
type AlitripMerchantGalaxyMemberRegisterDerbyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_member_register_derby_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 携带优惠券试单结果
	Result *AlitripMerchantGalaxyMemberRegisterDerbyResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyMemberRegisterDerbyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyMemberRegisterDerbyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyMemberRegisterDerbyAPIResponse)
	},
}

// GetAlitripMerchantGalaxyMemberRegisterDerbyAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyMemberRegisterDerbyAPIResponse
func GetAlitripMerchantGalaxyMemberRegisterDerbyAPIResponse() *AlitripMerchantGalaxyMemberRegisterDerbyAPIResponse {
	return poolAlitripMerchantGalaxyMemberRegisterDerbyAPIResponse.Get().(*AlitripMerchantGalaxyMemberRegisterDerbyAPIResponse)
}

// ReleaseAlitripMerchantGalaxyMemberRegisterDerbyAPIResponse 将 AlitripMerchantGalaxyMemberRegisterDerbyAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyMemberRegisterDerbyAPIResponse(v *AlitripMerchantGalaxyMemberRegisterDerbyAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyMemberRegisterDerbyAPIResponse.Put(v)
}
