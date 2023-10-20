package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopLsydListupoutAPIResponse 零售药店查询本企业上游企业出库单据信息 API返回值
// alibaba.alihealth.drugtrace.top.lsyd.listupout
//
// 查询货主/本企业上游企业出库单据信息
type AlibabaAlihealthDrugtraceTopLsydListupoutAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugtraceTopLsydListupoutAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopLsydListupoutAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugtraceTopLsydListupoutAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugtraceTopLsydListupoutAPIResponseModel is 零售药店查询本企业上游企业出库单据信息 成功返回结果
type AlibabaAlihealthDrugtraceTopLsydListupoutAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_lsyd_listupout_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugtraceTopLsydListupoutResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopLsydListupoutAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugtraceTopLsydListupoutAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopLsydListupoutAPIResponse)
	},
}

// GetAlibabaAlihealthDrugtraceTopLsydListupoutAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopLsydListupoutAPIResponse
func GetAlibabaAlihealthDrugtraceTopLsydListupoutAPIResponse() *AlibabaAlihealthDrugtraceTopLsydListupoutAPIResponse {
	return poolAlibabaAlihealthDrugtraceTopLsydListupoutAPIResponse.Get().(*AlibabaAlihealthDrugtraceTopLsydListupoutAPIResponse)
}

// ReleaseAlibabaAlihealthDrugtraceTopLsydListupoutAPIResponse 将 AlibabaAlihealthDrugtraceTopLsydListupoutAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopLsydListupoutAPIResponse(v *AlibabaAlihealthDrugtraceTopLsydListupoutAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopLsydListupoutAPIResponse.Put(v)
}
