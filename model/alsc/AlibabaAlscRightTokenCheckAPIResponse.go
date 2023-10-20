package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalscrighttokencheckAPIResponse 实物奖品凭证校验 API返回值
// alibaba.alsc.right.token.check
//
// 实物奖品凭证校验
type AlibabaalscrighttokencheckAPIResponse struct {
	model.CommonResponse
	AlibabaalscrighttokencheckAPIResponseModel
}

// AlibabaalscrighttokencheckAPIResponseModel is 实物奖品凭证校验 成功返回结果
type AlibabaalscrighttokencheckAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_right_token_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *BaseResponse `json:"result,omitempty" xml:"result,omitempty"`
}
