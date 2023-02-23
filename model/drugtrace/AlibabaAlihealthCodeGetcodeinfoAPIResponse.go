package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthCodeGetcodeinfoAPIResponse 码查询功能 API返回值
// alibaba.alihealth.code.getcodeinfo
//
// 码查询功能
type AlibabaAlihealthCodeGetcodeinfoAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthCodeGetcodeinfoAPIResponseModel
}

// AlibabaAlihealthCodeGetcodeinfoAPIResponseModel is 码查询功能 成功返回结果
type AlibabaAlihealthCodeGetcodeinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_code_getcodeinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口调用状态
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 接口调用状态码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 码信息类
	Model *DrugEntUseDto `json:"model,omitempty" xml:"model,omitempty"`
}
