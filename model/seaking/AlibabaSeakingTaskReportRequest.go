package seaking

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
跳转任务发布成功商品ID回传 API请求
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

// 初始化AlibabaSeakingTaskReportRequest对象
func NewAlibabaSeakingTaskReportRequest() *AlibabaSeakingTaskReportRequest{
    return &AlibabaSeakingTaskReportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSeakingTaskReportRequest) GetApiMethodName() string {
    return "alibaba.seaking.task.report"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSeakingTaskReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReportDetail Setter
// 上报数据详情
func (r *AlibabaSeakingTaskReportRequest) SetReportDetail(reportDetail []TaskDetailReportDto) error {
    r.reportDetail = reportDetail
    r.Set("report_detail", reportDetail)
    return nil
}

// ReportDetail Getter
func (r AlibabaSeakingTaskReportRequest) GetReportDetail() []TaskDetailReportDto {
    return r.reportDetail
}
// TaskType Setter
// 任务类型(title/image)
func (r *AlibabaSeakingTaskReportRequest) SetTaskType(taskType string) error {
    r.taskType = taskType
    r.Set("task_type", taskType)
    return nil
}

// TaskType Getter
func (r AlibabaSeakingTaskReportRequest) GetTaskType() string {
    return r.taskType
}
// Token Setter
// 用户token
func (r *AlibabaSeakingTaskReportRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r AlibabaSeakingTaskReportRequest) GetToken() string {
    return r.token
}
// TokenFrom Setter
// token来源站点
func (r *AlibabaSeakingTaskReportRequest) SetTokenFrom(tokenFrom string) error {
    r.tokenFrom = tokenFrom
    r.Set("token_from", tokenFrom)
    return nil
}

// TokenFrom Getter
func (r AlibabaSeakingTaskReportRequest) GetTokenFrom() string {
    return r.tokenFrom
}
