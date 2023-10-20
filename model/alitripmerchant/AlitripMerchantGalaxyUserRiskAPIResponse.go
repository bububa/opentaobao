package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyuserriskAPIResponse 查询微信账号的风险等级 API返回值
// alitrip.merchant.galaxy.user.risk
//
// 【星河服务】查询微信账号的风险等级
type AlitripmerchantgalaxyuserriskAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxyuserriskAPIResponseModel
}

// AlitripmerchantgalaxyuserriskAPIResponseModel is 查询微信账号的风险等级 成功返回结果
type AlitripmerchantgalaxyuserriskAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_user_risk_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripmerchantgalaxyuserriskResponse `json:"result,omitempty" xml:"result,omitempty"`
}
