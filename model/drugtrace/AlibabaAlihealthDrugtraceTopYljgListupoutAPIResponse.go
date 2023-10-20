package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopYljgListupoutAPIResponse 医疗机构查询本企业上游企业出库单据信息 API返回值
// alibaba.alihealth.drugtrace.top.yljg.listupout
//
// 查询货主/本企业上游企业出库单据信息
type AlibabaAlihealthDrugtraceTopYljgListupoutAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugtraceTopYljgListupoutAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopYljgListupoutAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugtraceTopYljgListupoutAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugtraceTopYljgListupoutAPIResponseModel is 医疗机构查询本企业上游企业出库单据信息 成功返回结果
type AlibabaAlihealthDrugtraceTopYljgListupoutAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_yljg_listupout_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugtraceTopYljgListupoutResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopYljgListupoutAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugtraceTopYljgListupoutAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopYljgListupoutAPIResponse)
	},
}

// GetAlibabaAlihealthDrugtraceTopYljgListupoutAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopYljgListupoutAPIResponse
func GetAlibabaAlihealthDrugtraceTopYljgListupoutAPIResponse() *AlibabaAlihealthDrugtraceTopYljgListupoutAPIResponse {
	return poolAlibabaAlihealthDrugtraceTopYljgListupoutAPIResponse.Get().(*AlibabaAlihealthDrugtraceTopYljgListupoutAPIResponse)
}

// ReleaseAlibabaAlihealthDrugtraceTopYljgListupoutAPIResponse 将 AlibabaAlihealthDrugtraceTopYljgListupoutAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopYljgListupoutAPIResponse(v *AlibabaAlihealthDrugtraceTopYljgListupoutAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopYljgListupoutAPIResponse.Put(v)
}
