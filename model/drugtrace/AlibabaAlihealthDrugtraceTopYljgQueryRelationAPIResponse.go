package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopYljgQueryRelationAPIResponse 单码关联关系查询 API返回值
// alibaba.alihealth.drugtrace.top.yljg.query.relation
//
// 单码关联关系查询
type AlibabaAlihealthDrugtraceTopYljgQueryRelationAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugtraceTopYljgQueryRelationAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopYljgQueryRelationAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugtraceTopYljgQueryRelationAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugtraceTopYljgQueryRelationAPIResponseModel is 单码关联关系查询 成功返回结果
type AlibabaAlihealthDrugtraceTopYljgQueryRelationAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_yljg_query_relation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihealthDrugtraceTopYljgQueryRelationResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopYljgQueryRelationAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugtraceTopYljgQueryRelationAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopYljgQueryRelationAPIResponse)
	},
}

// GetAlibabaAlihealthDrugtraceTopYljgQueryRelationAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopYljgQueryRelationAPIResponse
func GetAlibabaAlihealthDrugtraceTopYljgQueryRelationAPIResponse() *AlibabaAlihealthDrugtraceTopYljgQueryRelationAPIResponse {
	return poolAlibabaAlihealthDrugtraceTopYljgQueryRelationAPIResponse.Get().(*AlibabaAlihealthDrugtraceTopYljgQueryRelationAPIResponse)
}

// ReleaseAlibabaAlihealthDrugtraceTopYljgQueryRelationAPIResponse 将 AlibabaAlihealthDrugtraceTopYljgQueryRelationAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopYljgQueryRelationAPIResponse(v *AlibabaAlihealthDrugtraceTopYljgQueryRelationAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopYljgQueryRelationAPIResponse.Put(v)
}
