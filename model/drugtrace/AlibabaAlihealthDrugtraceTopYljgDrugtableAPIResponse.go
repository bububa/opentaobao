package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopYljgDrugtableAPIResponse 查询药品目录信息 API返回值
// alibaba.alihealth.drugtrace.top.yljg.drugtable
//
// 查询药品目录信息
type AlibabaAlihealthDrugtraceTopYljgDrugtableAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugtraceTopYljgDrugtableAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopYljgDrugtableAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugtraceTopYljgDrugtableAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugtraceTopYljgDrugtableAPIResponseModel is 查询药品目录信息 成功返回结果
type AlibabaAlihealthDrugtraceTopYljgDrugtableAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_yljg_drugtable_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugtraceTopYljgDrugtableResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopYljgDrugtableAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugtraceTopYljgDrugtableAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopYljgDrugtableAPIResponse)
	},
}

// GetAlibabaAlihealthDrugtraceTopYljgDrugtableAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopYljgDrugtableAPIResponse
func GetAlibabaAlihealthDrugtraceTopYljgDrugtableAPIResponse() *AlibabaAlihealthDrugtraceTopYljgDrugtableAPIResponse {
	return poolAlibabaAlihealthDrugtraceTopYljgDrugtableAPIResponse.Get().(*AlibabaAlihealthDrugtraceTopYljgDrugtableAPIResponse)
}

// ReleaseAlibabaAlihealthDrugtraceTopYljgDrugtableAPIResponse 将 AlibabaAlihealthDrugtraceTopYljgDrugtableAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopYljgDrugtableAPIResponse(v *AlibabaAlihealthDrugtraceTopYljgDrugtableAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopYljgDrugtableAPIResponse.Put(v)
}
