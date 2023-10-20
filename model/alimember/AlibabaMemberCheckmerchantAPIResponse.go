package alimember

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMemberCheckmerchantAPIResponse 校验商家身份 API返回值
// alibaba.member.checkmerchant
//
// 校验商家身份
type AlibabaMemberCheckmerchantAPIResponse struct {
	model.CommonResponse
	AlibabaMemberCheckmerchantAPIResponseModel
}

// AlibabaMemberCheckmerchantAPIResponseModel is 校验商家身份 成功返回结果
type AlibabaMemberCheckmerchantAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_member_checkmerchant_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// message
	ReturnMessage string `json:"return_message,omitempty" xml:"return_message,omitempty"`
	// code
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
}
