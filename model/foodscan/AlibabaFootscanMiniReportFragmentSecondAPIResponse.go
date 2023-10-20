package foodscan

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFootscanMiniReportFragmentSecondAPIResponse 第二只脚生成报告接口 API返回值
// alibaba.footscan.mini.report.fragment.second
//
// 第二只脚生成报告接口
type AlibabaFootscanMiniReportFragmentSecondAPIResponse struct {
	model.CommonResponse
	AlibabaFootscanMiniReportFragmentSecondAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaFootscanMiniReportFragmentSecondAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaFootscanMiniReportFragmentSecondAPIResponseModel).Reset()
}

// AlibabaFootscanMiniReportFragmentSecondAPIResponseModel is 第二只脚生成报告接口 成功返回结果
type AlibabaFootscanMiniReportFragmentSecondAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_footscan_mini_report_fragment_second_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务出参
	Result *AlibabaFootscanMiniReportFragmentSecondMtopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaFootscanMiniReportFragmentSecondAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaFootscanMiniReportFragmentSecondAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaFootscanMiniReportFragmentSecondAPIResponse)
	},
}

// GetAlibabaFootscanMiniReportFragmentSecondAPIResponse 从 sync.Pool 获取 AlibabaFootscanMiniReportFragmentSecondAPIResponse
func GetAlibabaFootscanMiniReportFragmentSecondAPIResponse() *AlibabaFootscanMiniReportFragmentSecondAPIResponse {
	return poolAlibabaFootscanMiniReportFragmentSecondAPIResponse.Get().(*AlibabaFootscanMiniReportFragmentSecondAPIResponse)
}

// ReleaseAlibabaFootscanMiniReportFragmentSecondAPIResponse 将 AlibabaFootscanMiniReportFragmentSecondAPIResponse 保存到 sync.Pool
func ReleaseAlibabaFootscanMiniReportFragmentSecondAPIResponse(v *AlibabaFootscanMiniReportFragmentSecondAPIResponse) {
	v.Reset()
	poolAlibabaFootscanMiniReportFragmentSecondAPIResponse.Put(v)
}
