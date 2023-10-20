package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthcodegetcodeinfoAPIResponse 码查询功能 API返回值
// alibaba.alihealth.code.getcodeinfo
//
// 码查询功能
type AlibabaalihealthcodegetcodeinfoAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthcodegetcodeinfoAPIResponseModel
}

// AlibabaalihealthcodegetcodeinfoAPIResponseModel is 码查询功能 成功返回结果
type AlibabaalihealthcodegetcodeinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_code_getcodeinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口调用状态，FAIL_BIZ_CODE_FORMAT_ERROR:追溯码格式错误,FAIL_BIZ_CODE_NO_AUTH:权限不足,FAIL_BIZ_CODE_NOT_EXIST:不是追溯码,FAIL_BIZ_CODE_NOT_ACTIVE:追溯码未激活,FAIL_BIZ_REMOTE_REQ_ERROR:调用远程服务失败,1000:调用成功
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 接口调用状态码，FAIL_BIZ_CODE_FORMAT_ERROR:追溯码格式错误,FAIL_BIZ_CODE_NO_AUTH:权限不足,FAIL_BIZ_CODE_NOT_EXIST:不是追溯码,FAIL_BIZ_CODE_NOT_ACTIVE:追溯码未激活,FAIL_BIZ_REMOTE_REQ_ERROR:调用远程服务失败,1000:调用成功
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 码信息类
	Model *DrugEntUseDto `json:"model,omitempty" xml:"model,omitempty"`
}
