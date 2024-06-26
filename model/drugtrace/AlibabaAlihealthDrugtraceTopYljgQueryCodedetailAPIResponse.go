package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIResponse 根据码查询码信息 API返回值
// alibaba.alihealth.drugtrace.top.yljg.query.codedetail
//
// 服务描述
// 此接口，针对有码药品，提供可通过追溯码获取该药品的基础信息和生产状况信息。
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
type AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIResponseModel is 根据码查询码信息 成功返回结果
type AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_yljg_query_codedetail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 最外层结果
	Result *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIResponse)
	},
}

// GetAlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIResponse
func GetAlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIResponse() *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIResponse {
	return poolAlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIResponse.Get().(*AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIResponse)
}

// ReleaseAlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIResponse 将 AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIResponse(v *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIResponse.Put(v)
}
