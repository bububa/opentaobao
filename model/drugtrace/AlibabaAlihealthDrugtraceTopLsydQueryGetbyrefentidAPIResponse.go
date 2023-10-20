package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIResponse 根据企业唯一标识查看企业详细信息 API返回值
// alibaba.alihealth.drugtrace.top.lsyd.query.getbyrefentid
//
// 根据企业唯一标识查看企业详细信息
type AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIResponseModel is 根据企业唯一标识查看企业详细信息 成功返回结果
type AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_lsyd_query_getbyrefentid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIResponse)
	},
}

// GetAlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIResponse
func GetAlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIResponse() *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIResponse {
	return poolAlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIResponse.Get().(*AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIResponse)
}

// ReleaseAlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIResponse 将 AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIResponse(v *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIResponse.Put(v)
}
