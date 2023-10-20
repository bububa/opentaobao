package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopYljgQueryBillstatusAPIResponse 上传单据后处理状态查询 API返回值
// alibaba.alihealth.drugtrace.top.yljg.query.billstatus
//
// 单据处理状态查询
type AlibabaAlihealthDrugtraceTopYljgQueryBillstatusAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugtraceTopYljgQueryBillstatusAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugtraceTopYljgQueryBillstatusAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugtraceTopYljgQueryBillstatusAPIResponseModel is 上传单据后处理状态查询 成功返回结果
type AlibabaAlihealthDrugtraceTopYljgQueryBillstatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_yljg_query_billstatus_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugtraceTopYljgQueryBillstatusAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopYljgQueryBillstatusAPIResponse)
	},
}

// GetAlibabaAlihealthDrugtraceTopYljgQueryBillstatusAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopYljgQueryBillstatusAPIResponse
func GetAlibabaAlihealthDrugtraceTopYljgQueryBillstatusAPIResponse() *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusAPIResponse {
	return poolAlibabaAlihealthDrugtraceTopYljgQueryBillstatusAPIResponse.Get().(*AlibabaAlihealthDrugtraceTopYljgQueryBillstatusAPIResponse)
}

// ReleaseAlibabaAlihealthDrugtraceTopYljgQueryBillstatusAPIResponse 将 AlibabaAlihealthDrugtraceTopYljgQueryBillstatusAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopYljgQueryBillstatusAPIResponse(v *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopYljgQueryBillstatusAPIResponse.Put(v)
}
