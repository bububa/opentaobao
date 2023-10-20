package security

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDiafiTokenCheckAPIResponse 天朗token校验API API返回值
// alibaba.diafi.token.check
//
// 天朗token校验
type AlibabaDiafiTokenCheckAPIResponse struct {
	model.CommonResponse
	AlibabaDiafiTokenCheckAPIResponseModel
}

// AlibabaDiafiTokenCheckAPIResponseModel is 天朗token校验API 成功返回结果
type AlibabaDiafiTokenCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_diafi_token_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果集
	Result *DiAfiResult `json:"result,omitempty" xml:"result,omitempty"`
}
