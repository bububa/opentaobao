package alitripmerchant

import (
	"encoding/xml"

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

// AlitripMerchantGalaxyMemberAddAgreementAPIResponseModel is 添加用户协议记录接口 成功返回结果
type AlitripMerchantGalaxyMemberAddAgreementAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_member_add_agreement_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlitripMerchantGalaxyMemberAddAgreementResponse `json:"result,omitempty" xml:"result,omitempty"`
}
