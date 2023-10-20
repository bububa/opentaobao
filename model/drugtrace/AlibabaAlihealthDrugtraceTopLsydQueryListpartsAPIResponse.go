package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIResponse 往来单位查询 API返回值
// alibaba.alihealth.drugtrace.top.lsyd.query.listparts
//
// 查询往来单位列表
type AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIResponseModel is 往来单位查询 成功返回结果
type AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_lsyd_query_listparts_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugtraceTopLsydQueryListpartsResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIResponse)
	},
}

// GetAlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIResponse
func GetAlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIResponse() *AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIResponse {
	return poolAlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIResponse.Get().(*AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIResponse)
}

// ReleaseAlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIResponse 将 AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIResponse(v *AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIResponse.Put(v)
}
