package seaking

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSeakingTaskReportAPIResponse 跳转任务发布成功商品ID回传 API返回值
// alibaba.seaking.task.report
//
// 跳转任务发布成功商品ID回传
type AlibabaSeakingTaskReportAPIResponse struct {
	model.CommonResponse
	AlibabaSeakingTaskReportAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSeakingTaskReportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSeakingTaskReportAPIResponseModel).Reset()
}

// AlibabaSeakingTaskReportAPIResponseModel is 跳转任务发布成功商品ID回传 成功返回结果
type AlibabaSeakingTaskReportAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_seaking_task_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSeakingTaskReportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolAlibabaSeakingTaskReportAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSeakingTaskReportAPIResponse)
	},
}

// GetAlibabaSeakingTaskReportAPIResponse 从 sync.Pool 获取 AlibabaSeakingTaskReportAPIResponse
func GetAlibabaSeakingTaskReportAPIResponse() *AlibabaSeakingTaskReportAPIResponse {
	return poolAlibabaSeakingTaskReportAPIResponse.Get().(*AlibabaSeakingTaskReportAPIResponse)
}

// ReleaseAlibabaSeakingTaskReportAPIResponse 将 AlibabaSeakingTaskReportAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSeakingTaskReportAPIResponse(v *AlibabaSeakingTaskReportAPIResponse) {
	v.Reset()
	poolAlibabaSeakingTaskReportAPIResponse.Put(v)
}
