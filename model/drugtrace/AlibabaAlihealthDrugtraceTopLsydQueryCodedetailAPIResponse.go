package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopLsydQueryCodedetailAPIResponse 根据码查询码信息 API返回值
// alibaba.alihealth.drugtrace.top.lsyd.query.codedetail
//
// 服务描述
// 此接口，针对有码药品，提供可通过追溯码获取该药品的基础信息和生产状况信息。
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
type AlibabaAlihealthDrugtraceTopLsydQueryCodedetailAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugtraceTopLsydQueryCodedetailAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugtraceTopLsydQueryCodedetailAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugtraceTopLsydQueryCodedetailAPIResponseModel is 根据码查询码信息 成功返回结果
type AlibabaAlihealthDrugtraceTopLsydQueryCodedetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_lsyd_query_codedetail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 最外层结果
	Result *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugtraceTopLsydQueryCodedetailAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopLsydQueryCodedetailAPIResponse)
	},
}

// GetAlibabaAlihealthDrugtraceTopLsydQueryCodedetailAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopLsydQueryCodedetailAPIResponse
func GetAlibabaAlihealthDrugtraceTopLsydQueryCodedetailAPIResponse() *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailAPIResponse {
	return poolAlibabaAlihealthDrugtraceTopLsydQueryCodedetailAPIResponse.Get().(*AlibabaAlihealthDrugtraceTopLsydQueryCodedetailAPIResponse)
}

// ReleaseAlibabaAlihealthDrugtraceTopLsydQueryCodedetailAPIResponse 将 AlibabaAlihealthDrugtraceTopLsydQueryCodedetailAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopLsydQueryCodedetailAPIResponse(v *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopLsydQueryCodedetailAPIResponse.Put(v)
}
