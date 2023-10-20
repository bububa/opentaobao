package charity

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacharitycharitytimecommonauthAPIResponse 通用授权 API返回值
// alibaba.charity.charitytime.commonauth
//
// 三小时和外部账号绑定通用top 返回跳转链接进行绑定
type AlibabacharitycharitytimecommonauthAPIResponse struct {
	model.CommonResponse
	AlibabacharitycharitytimecommonauthAPIResponseModel
}

// AlibabacharitycharitytimecommonauthAPIResponseModel is 通用授权 成功返回结果
type AlibabacharitycharitytimecommonauthAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_charity_charitytime_commonauth_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *CsrResult `json:"result,omitempty" xml:"result,omitempty"`
}
