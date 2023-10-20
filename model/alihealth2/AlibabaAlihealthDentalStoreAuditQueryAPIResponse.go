package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDentalStoreAuditQueryAPIResponse ISV查询门店审核状态 API返回值
// alibaba.alihealth.dental.store.audit.query
//
// ISV查询门店审核状态
type AlibabaAlihealthDentalStoreAuditQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDentalStoreAuditQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDentalStoreAuditQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDentalStoreAuditQueryAPIResponseModel).Reset()
}

// AlibabaAlihealthDentalStoreAuditQueryAPIResponseModel is ISV查询门店审核状态 成功返回结果
type AlibabaAlihealthDentalStoreAuditQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_dental_store_audit_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAlihealthDentalStoreAuditQueryMtopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDentalStoreAuditQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDentalStoreAuditQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDentalStoreAuditQueryAPIResponse)
	},
}

// GetAlibabaAlihealthDentalStoreAuditQueryAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDentalStoreAuditQueryAPIResponse
func GetAlibabaAlihealthDentalStoreAuditQueryAPIResponse() *AlibabaAlihealthDentalStoreAuditQueryAPIResponse {
	return poolAlibabaAlihealthDentalStoreAuditQueryAPIResponse.Get().(*AlibabaAlihealthDentalStoreAuditQueryAPIResponse)
}

// ReleaseAlibabaAlihealthDentalStoreAuditQueryAPIResponse 将 AlibabaAlihealthDentalStoreAuditQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDentalStoreAuditQueryAPIResponse(v *AlibabaAlihealthDentalStoreAuditQueryAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDentalStoreAuditQueryAPIResponse.Put(v)
}
