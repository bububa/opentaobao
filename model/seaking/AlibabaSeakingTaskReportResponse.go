package seaking

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
跳转任务发布成功商品ID回传 APIResponse
alibaba.seaking.task.report

跳转任务发布成功商品ID回传
*/
type AlibabaSeakingTaskReportAPIResponse struct {
    model.CommonResponse
    AlibabaSeakingTaskReportResponse
}

type AlibabaSeakingTaskReportResponse struct {
    XMLName xml.Name `xml:"alibaba_seaking_task_report_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 系统自动生成
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
