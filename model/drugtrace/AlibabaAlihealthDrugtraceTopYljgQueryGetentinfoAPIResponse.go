package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIResponse 通过企业名得到唯一标识ref_ent_id及企业ent_id API返回值
// alibaba.alihealth.drugtrace.top.yljg.query.getentinfo
//
// 根据企业名称查询ID
type AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIResponseModel is 通过企业名得到唯一标识ref_ent_id及企业ent_id 成功返回结果
type AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_yljg_query_getentinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIResponse)
	},
}

// GetAlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIResponse
func GetAlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIResponse() *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIResponse {
	return poolAlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIResponse.Get().(*AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIResponse)
}

// ReleaseAlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIResponse 将 AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIResponse(v *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIResponse.Put(v)
}
