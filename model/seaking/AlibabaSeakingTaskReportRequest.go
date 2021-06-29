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
    _reportDetail   []TaskDetailReportDto
    // 任务类型(title/image)
    _taskType   string
    // 用户token
    _token   string
    // token来源站点
    _tokenFrom   string
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
func (r *AlibabaSeakingTaskReportRequest) SetReportDetail(_reportDetail []TaskDetailReportDto) error {
    r._reportDetail = _reportDetail
    r.Set("report_detail", _reportDetail)
    return nil
}

// ReportDetail Getter
func (r AlibabaSeakingTaskReportRequest) GetReportDetail() []TaskDetailReportDto {
    return r._reportDetail
}
// TaskType Setter
// 任务类型(title/image)
func (r *AlibabaSeakingTaskReportRequest) SetTaskType(_taskType string) error {
    r._taskType = _taskType
    r.Set("task_type", _taskType)
    return nil
}

// TaskType Getter
func (r AlibabaSeakingTaskReportRequest) GetTaskType() string {
    return r._taskType
}
// Token Setter
// 用户token
func (r *AlibabaSeakingTaskReportRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r AlibabaSeakingTaskReportRequest) GetToken() string {
    return r._token
}
// TokenFrom Setter
// token来源站点
func (r *AlibabaSeakingTaskReportRequest) SetTokenFrom(_tokenFrom string) error {
    r._tokenFrom = _tokenFrom
    r.Set("token_from", _tokenFrom)
    return nil
}

// TokenFrom Getter
func (r AlibabaSeakingTaskReportRequest) GetTokenFrom() string {
    return r._tokenFrom
}
