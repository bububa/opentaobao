package seaking

import (
	"encoding/xml"

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

// AlibabaSeakingTaskReportAPIResponseModel is 跳转任务发布成功商品ID回传 成功返回结果
type AlibabaSeakingTaskReportAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_seaking_task_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
