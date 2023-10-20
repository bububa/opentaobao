package alihealthpw

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPwGmAuditAPIResponse 同情用药审核接口 API返回值
// alibaba.alihealth.pw.gm.audit
//
// 同情用药审核接口，提供给合作方审核申请单
type AlibabaAlihealthPwGmAuditAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthPwGmAuditAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthPwGmAuditAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthPwGmAuditAPIResponseModel).Reset()
}

// AlibabaAlihealthPwGmAuditAPIResponseModel is 同情用药审核接口 成功返回结果
type AlibabaAlihealthPwGmAuditAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_pw_gm_audit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *ResponseMessage `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthPwGmAuditAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthPwGmAuditAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthPwGmAuditAPIResponse)
	},
}

// GetAlibabaAlihealthPwGmAuditAPIResponse 从 sync.Pool 获取 AlibabaAlihealthPwGmAuditAPIResponse
func GetAlibabaAlihealthPwGmAuditAPIResponse() *AlibabaAlihealthPwGmAuditAPIResponse {
	return poolAlibabaAlihealthPwGmAuditAPIResponse.Get().(*AlibabaAlihealthPwGmAuditAPIResponse)
}

// ReleaseAlibabaAlihealthPwGmAuditAPIResponse 将 AlibabaAlihealthPwGmAuditAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthPwGmAuditAPIResponse(v *AlibabaAlihealthPwGmAuditAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthPwGmAuditAPIResponse.Put(v)
}
