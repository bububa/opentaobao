package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIResponse 根据企业唯一标识查看企业详细信息 API返回值
// alibaba.alihealth.drugtrace.top.yljg.query.getbyrefentid
//
// 根据企业唯一标识查看企业详细信息
type AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIResponseModel is 根据企业唯一标识查看企业详细信息 成功返回结果
type AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_yljg_query_getbyrefentid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIResponse)
	},
}

// GetAlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIResponse
func GetAlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIResponse() *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIResponse {
	return poolAlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIResponse.Get().(*AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIResponse)
}

// ReleaseAlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIResponse 将 AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIResponse(v *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIResponse.Put(v)
}
