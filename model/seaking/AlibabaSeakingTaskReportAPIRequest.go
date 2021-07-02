package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSeakingTaskReportAPIRequest 跳转任务发布成功商品ID回传 API请求
// alibaba.seaking.task.report
//
// 跳转任务发布成功商品ID回传
type AlibabaSeakingTaskReportAPIRequest struct {
	model.Params
	// 上报数据详情
	_reportDetail []TaskDetailReportDto
	// 任务类型(title/image)
	_taskType string
	// 用户token
	_token string
	// token来源站点
	_tokenFrom string
}

// NewAlibabaSeakingTaskReportRequest 初始化AlibabaSeakingTaskReportAPIRequest对象
func NewAlibabaSeakingTaskReportRequest() *AlibabaSeakingTaskReportAPIRequest {
	return &AlibabaSeakingTaskReportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSeakingTaskReportAPIRequest) GetApiMethodName() string {
	return "alibaba.seaking.task.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSeakingTaskReportAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ReportDetail Setter
// 上报数据详情
func (r *AlibabaSeakingTaskReportAPIRequest) SetReportDetail(_reportDetail []TaskDetailReportDto) error {
	r._reportDetail = _reportDetail
	r.Set("report_detail", _reportDetail)
	return nil
}

// Get ReportDetail Getter
func (r AlibabaSeakingTaskReportAPIRequest) GetReportDetail() []TaskDetailReportDto {
	return r._reportDetail
}

// Set is TaskType Setter
// 任务类型(title/image)
func (r *AlibabaSeakingTaskReportAPIRequest) SetTaskType(_taskType string) error {
	r._taskType = _taskType
	r.Set("task_type", _taskType)
	return nil
}

// Get TaskType Getter
func (r AlibabaSeakingTaskReportAPIRequest) GetTaskType() string {
	return r._taskType
}

// Set is Token Setter
// 用户token
func (r *AlibabaSeakingTaskReportAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// Get Token Getter
func (r AlibabaSeakingTaskReportAPIRequest) GetToken() string {
	return r._token
}

// Set is TokenFrom Setter
// token来源站点
func (r *AlibabaSeakingTaskReportAPIRequest) SetTokenFrom(_tokenFrom string) error {
	r._tokenFrom = _tokenFrom
	r.Set("token_from", _tokenFrom)
	return nil
}

// Get TokenFrom Getter
func (r AlibabaSeakingTaskReportAPIRequest) GetTokenFrom() string {
	return r._tokenFrom
}
