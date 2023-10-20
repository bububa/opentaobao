package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIResponse 上游出库单单据明细查询 API返回值
// alibaba.alihealth.drugtrace.top.lsyd.listupout.detail
//
// 查询上游出库单明细(带追溯码信息)
type AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIResponseModel is 上游出库单单据明细查询 成功返回结果
type AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_lsyd_listupout_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugtraceTopLsydListupoutDetailResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIResponse)
	},
}

// GetAlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIResponse
func GetAlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIResponse() *AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIResponse {
	return poolAlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIResponse.Get().(*AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIResponse)
}

// ReleaseAlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIResponse 将 AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIResponse(v *AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIResponse.Put(v)
}
