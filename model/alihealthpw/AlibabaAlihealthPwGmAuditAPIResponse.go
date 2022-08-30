package alihealthpw

import (
	"encoding/xml"

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

// AlibabaAlihealthPwGmAuditAPIResponseModel is 同情用药审核接口 成功返回结果
type AlibabaAlihealthPwGmAuditAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_pw_gm_audit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *ResponseMessage `json:"result,omitempty" xml:"result,omitempty"`
}
