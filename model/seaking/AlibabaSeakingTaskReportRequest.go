package seaking

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
跳转任务发布成功商品ID回传 APIRequest
alibaba.seaking.task.report

跳转任务发布成功商品ID回传
*/
type AlibabaSeakingTaskReportRequest struct {
    model.Params

    // 上报数据详情
    reportDetail   []TaskDetailReportDto 

    // 任务类型(title/image)
    taskType   string 

    // 用户token
    token   string 

    // token来源站点
    tokenFrom   string 

}

func NewAlibabaSeakingTaskReportRequest() *AlibabaSeakingTaskReportRequest{
    return &AlibabaSeakingTaskReportRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSeakingTaskReportRequest) GetApiMethodName() string {
    return "alibaba.seaking.task.report"
}

func (r AlibabaSeakingTaskReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSeakingTaskReportRequest) SetReportDetail(reportDetail []TaskDetailReportDto) error {
    r.reportDetail = reportDetail
    r.Set("report_detail", reportDetail)
    return nil
}

func (r AlibabaSeakingTaskReportRequest) GetReportDetail() []TaskDetailReportDto {
    return r.reportDetail
}

func (r *AlibabaSeakingTaskReportRequest) SetTaskType(taskType string) error {
    r.taskType = taskType
    r.Set("task_type", taskType)
    return nil
}

func (r AlibabaSeakingTaskReportRequest) GetTaskType() string {
    return r.taskType
}

func (r *AlibabaSeakingTaskReportRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r AlibabaSeakingTaskReportRequest) GetToken() string {
    return r.token
}

func (r *AlibabaSeakingTaskReportRequest) SetTokenFrom(tokenFrom string) error {
    r.tokenFrom = tokenFrom
    r.Set("token_from", tokenFrom)
    return nil
}

func (r AlibabaSeakingTaskReportRequest) GetTokenFrom() string {
    return r.tokenFrom
}

