package foodscan

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFootscanMiniReportFragmentFirstAPIResponse 第一只脚生成报告接口 API返回值
// alibaba.footscan.mini.report.fragment.first
//
// 第一只脚生成报告接口
type AlibabaFootscanMiniReportFragmentFirstAPIResponse struct {
	model.CommonResponse
	AlibabaFootscanMiniReportFragmentFirstAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaFootscanMiniReportFragmentFirstAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaFootscanMiniReportFragmentFirstAPIResponseModel).Reset()
}

// AlibabaFootscanMiniReportFragmentFirstAPIResponseModel is 第一只脚生成报告接口 成功返回结果
type AlibabaFootscanMiniReportFragmentFirstAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_footscan_mini_report_fragment_first_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务出参
	Result *AlibabaFootscanMiniReportFragmentFirstMtopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaFootscanMiniReportFragmentFirstAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaFootscanMiniReportFragmentFirstAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaFootscanMiniReportFragmentFirstAPIResponse)
	},
}

// GetAlibabaFootscanMiniReportFragmentFirstAPIResponse 从 sync.Pool 获取 AlibabaFootscanMiniReportFragmentFirstAPIResponse
func GetAlibabaFootscanMiniReportFragmentFirstAPIResponse() *AlibabaFootscanMiniReportFragmentFirstAPIResponse {
	return poolAlibabaFootscanMiniReportFragmentFirstAPIResponse.Get().(*AlibabaFootscanMiniReportFragmentFirstAPIResponse)
}

// ReleaseAlibabaFootscanMiniReportFragmentFirstAPIResponse 将 AlibabaFootscanMiniReportFragmentFirstAPIResponse 保存到 sync.Pool
func ReleaseAlibabaFootscanMiniReportFragmentFirstAPIResponse(v *AlibabaFootscanMiniReportFragmentFirstAPIResponse) {
	v.Reset()
	poolAlibabaFootscanMiniReportFragmentFirstAPIResponse.Put(v)
}
