package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyMemberPopupAgreementAPIResponse 小程序唤起协议弹窗 API返回值
// alitrip.merchant.galaxy.member.popup.agreement
//
// 用户进入小程序后，根据用户是否授权协议，判断是否唤起协议弹窗
type AlitripMerchantGalaxyMemberPopupAgreementAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyMemberPopupAgreementAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyMemberPopupAgreementAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyMemberPopupAgreementAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyMemberPopupAgreementAPIResponseModel is 小程序唤起协议弹窗 成功返回结果
type AlitripMerchantGalaxyMemberPopupAgreementAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_member_popup_agreement_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果集
	Result *AlitripMerchantGalaxyMemberPopupAgreementResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyMemberPopupAgreementAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyMemberPopupAgreementAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyMemberPopupAgreementAPIResponse)
	},
}

// GetAlitripMerchantGalaxyMemberPopupAgreementAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyMemberPopupAgreementAPIResponse
func GetAlitripMerchantGalaxyMemberPopupAgreementAPIResponse() *AlitripMerchantGalaxyMemberPopupAgreementAPIResponse {
	return poolAlitripMerchantGalaxyMemberPopupAgreementAPIResponse.Get().(*AlitripMerchantGalaxyMemberPopupAgreementAPIResponse)
}

// ReleaseAlitripMerchantGalaxyMemberPopupAgreementAPIResponse 将 AlitripMerchantGalaxyMemberPopupAgreementAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyMemberPopupAgreementAPIResponse(v *AlitripMerchantGalaxyMemberPopupAgreementAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyMemberPopupAgreementAPIResponse.Put(v)
}
