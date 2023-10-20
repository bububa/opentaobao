package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxymemberpopupagreementAPIResponse 小程序唤起协议弹窗 API返回值
// alitrip.merchant.galaxy.member.popup.agreement
//
// 用户进入小程序后，根据用户是否授权协议，判断是否唤起协议弹窗
type AlitripmerchantgalaxymemberpopupagreementAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxymemberpopupagreementAPIResponseModel
}

// AlitripmerchantgalaxymemberpopupagreementAPIResponseModel is 小程序唤起协议弹窗 成功返回结果
type AlitripmerchantgalaxymemberpopupagreementAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_member_popup_agreement_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果集
	Result *AlitripmerchantgalaxymemberpopupagreementResponse `json:"result,omitempty" xml:"result,omitempty"`
}
