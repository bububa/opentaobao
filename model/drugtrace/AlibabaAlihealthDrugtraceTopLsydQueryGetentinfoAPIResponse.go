package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIResponse 通过企业名得到唯一标识ref_ent_id及企业ent_id API返回值
// alibaba.alihealth.drugtrace.top.lsyd.query.getentinfo
//
// 根据企业名称查询ID
type AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIResponseModel is 通过企业名得到唯一标识ref_ent_id及企业ent_id 成功返回结果
type AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_lsyd_query_getentinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIResponse)
	},
}

// GetAlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIResponse
func GetAlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIResponse() *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIResponse {
	return poolAlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIResponse.Get().(*AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIResponse)
}

// ReleaseAlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIResponse 将 AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIResponse(v *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIResponse.Put(v)
}
