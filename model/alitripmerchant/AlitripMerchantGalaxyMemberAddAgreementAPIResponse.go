package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyMemberAddAgreementAPIResponse 添加用户协议记录接口 API返回值
// alitrip.merchant.galaxy.member.add.agreement
//
// 记录用户是否授权协议
type AlitripMerchantGalaxyMemberAddAgreementAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyMemberAddAgreementAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyMemberAddAgreementAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyMemberAddAgreementAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyMemberAddAgreementAPIResponseModel is 添加用户协议记录接口 成功返回结果
type AlitripMerchantGalaxyMemberAddAgreementAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_member_add_agreement_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlitripMerchantGalaxyMemberAddAgreementResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyMemberAddAgreementAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyMemberAddAgreementAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyMemberAddAgreementAPIResponse)
	},
}

// GetAlitripMerchantGalaxyMemberAddAgreementAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyMemberAddAgreementAPIResponse
func GetAlitripMerchantGalaxyMemberAddAgreementAPIResponse() *AlitripMerchantGalaxyMemberAddAgreementAPIResponse {
	return poolAlitripMerchantGalaxyMemberAddAgreementAPIResponse.Get().(*AlitripMerchantGalaxyMemberAddAgreementAPIResponse)
}

// ReleaseAlitripMerchantGalaxyMemberAddAgreementAPIResponse 将 AlitripMerchantGalaxyMemberAddAgreementAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyMemberAddAgreementAPIResponse(v *AlitripMerchantGalaxyMemberAddAgreementAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyMemberAddAgreementAPIResponse.Put(v)
}
