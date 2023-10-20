package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIResponse 非药单码替换 API返回值
// alibaba.alihealth.tracecodeseller.code.single.codereplace
//
// 提供非药追溯码单码替换功能
type AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIResponseModel
}

// AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIResponseModel is 非药单码替换 成功返回结果
type AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_tracecodeseller_code_single_codereplace_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 操作说明
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 操作码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
}
