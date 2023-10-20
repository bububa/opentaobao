package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxymemberregisterderbyAPIResponse 会员注册(新版注册接口对接德比) API返回值
// alitrip.merchant.galaxy.member.register.derby
//
// 会员注册(新版注册接口对接德比)，返回手机号/邮箱/注册状态
type AlitripmerchantgalaxymemberregisterderbyAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxymemberregisterderbyAPIResponseModel
}

// AlitripmerchantgalaxymemberregisterderbyAPIResponseModel is 会员注册(新版注册接口对接德比) 成功返回结果
type AlitripmerchantgalaxymemberregisterderbyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_member_register_derby_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 携带优惠券试单结果
	Result *AlitripmerchantgalaxymemberregisterderbyResponse `json:"result,omitempty" xml:"result,omitempty"`
}
