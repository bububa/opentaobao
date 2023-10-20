package drugtrace

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaAlihealthCodeGetcodeinfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthCodeGetcodeinfoAPIResponseModel).Reset()
}

// AlibabaAlihealthCodeGetcodeinfoAPIResponseModel is 码查询功能 成功返回结果
type AlibabaAlihealthCodeGetcodeinfoAPIResponseModel struct {
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

// Reset 清空结构体
func (m *AlibabaAlihealthCodeGetcodeinfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.Model = nil
}

var poolAlibabaAlihealthCodeGetcodeinfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthCodeGetcodeinfoAPIResponse)
	},
}

// GetAlibabaAlihealthCodeGetcodeinfoAPIResponse 从 sync.Pool 获取 AlibabaAlihealthCodeGetcodeinfoAPIResponse
func GetAlibabaAlihealthCodeGetcodeinfoAPIResponse() *AlibabaAlihealthCodeGetcodeinfoAPIResponse {
	return poolAlibabaAlihealthCodeGetcodeinfoAPIResponse.Get().(*AlibabaAlihealthCodeGetcodeinfoAPIResponse)
}

// ReleaseAlibabaAlihealthCodeGetcodeinfoAPIResponse 将 AlibabaAlihealthCodeGetcodeinfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthCodeGetcodeinfoAPIResponse(v *AlibabaAlihealthCodeGetcodeinfoAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthCodeGetcodeinfoAPIResponse.Put(v)
}
