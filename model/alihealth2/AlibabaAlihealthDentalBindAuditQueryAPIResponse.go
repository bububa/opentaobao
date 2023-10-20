package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDentalBindAuditQueryAPIResponse ISV查询绑定审核状态 API返回值
// alibaba.alihealth.dental.bind.audit.query
//
// ISV查询绑定审核状态
type AlibabaAlihealthDentalBindAuditQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDentalBindAuditQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDentalBindAuditQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDentalBindAuditQueryAPIResponseModel).Reset()
}

// AlibabaAlihealthDentalBindAuditQueryAPIResponseModel is ISV查询绑定审核状态 成功返回结果
type AlibabaAlihealthDentalBindAuditQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_dental_bind_audit_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAlihealthDentalBindAuditQueryMtopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDentalBindAuditQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDentalBindAuditQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDentalBindAuditQueryAPIResponse)
	},
}

// GetAlibabaAlihealthDentalBindAuditQueryAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDentalBindAuditQueryAPIResponse
func GetAlibabaAlihealthDentalBindAuditQueryAPIResponse() *AlibabaAlihealthDentalBindAuditQueryAPIResponse {
	return poolAlibabaAlihealthDentalBindAuditQueryAPIResponse.Get().(*AlibabaAlihealthDentalBindAuditQueryAPIResponse)
}

// ReleaseAlibabaAlihealthDentalBindAuditQueryAPIResponse 将 AlibabaAlihealthDentalBindAuditQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDentalBindAuditQueryAPIResponse(v *AlibabaAlihealthDentalBindAuditQueryAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDentalBindAuditQueryAPIResponse.Put(v)
}
