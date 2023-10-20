package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIResponse 单码关联关系查询 API返回值
// alibaba.alihealth.drugtrace.top.lsyd.query.relation
//
// 单码关联关系查询
type AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIResponseModel is 单码关联关系查询 成功返回结果
type AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_lsyd_query_relation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihealthDrugtraceTopLsydQueryRelationResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugtraceTopLsydQueryRelationAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIResponse)
	},
}

// GetAlibabaAlihealthDrugtraceTopLsydQueryRelationAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIResponse
func GetAlibabaAlihealthDrugtraceTopLsydQueryRelationAPIResponse() *AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIResponse {
	return poolAlibabaAlihealthDrugtraceTopLsydQueryRelationAPIResponse.Get().(*AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIResponse)
}

// ReleaseAlibabaAlihealthDrugtraceTopLsydQueryRelationAPIResponse 将 AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopLsydQueryRelationAPIResponse(v *AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopLsydQueryRelationAPIResponse.Put(v)
}
