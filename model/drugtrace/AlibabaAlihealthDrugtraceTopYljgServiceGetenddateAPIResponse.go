package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIResponse 获取服务截止日期 API返回值
// alibaba.alihealth.drugtrace.top.yljg.service.getenddate
//
// 获取企业服务截止时间
type AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIResponseModel is 获取服务截止日期 成功返回结果
type AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_yljg_service_getenddate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回
	Result *AlibabaAlihealthDrugtraceTopYljgServiceGetenddateResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIResponse)
	},
}

// GetAlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIResponse
func GetAlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIResponse() *AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIResponse {
	return poolAlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIResponse.Get().(*AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIResponse)
}

// ReleaseAlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIResponse 将 AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIResponse(v *AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIResponse.Put(v)
}
