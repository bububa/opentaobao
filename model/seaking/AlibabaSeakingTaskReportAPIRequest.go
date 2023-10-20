package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaseakingtaskreportAPIRequest 跳转任务发布成功商品ID回传 API请求
// alibaba.seaking.task.report
//
// 跳转任务发布成功商品ID回传
type AlibabaseakingtaskreportAPIRequest struct {
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

// NewAlibabaseakingtaskreportRequest 初始化AlibabaseakingtaskreportAPIRequest对象
func NewAlibabaseakingtaskreportRequest() *AlibabaseakingtaskreportAPIRequest {
	return &AlibabaseakingtaskreportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaseakingtaskreportAPIRequest) GetApiMethodName() string {
	return "alibaba.seaking.task.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaseakingtaskreportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaseakingtaskreportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReportDetail is ReportDetail Setter
// 上报数据详情
func (r *AlibabaseakingtaskreportAPIRequest) SetReportDetail(_reportDetail []TaskDetailReportDto) error {
	r._reportDetail = _reportDetail
	r.Set("report_detail", _reportDetail)
	return nil
}

// GetReportDetail ReportDetail Getter
func (r AlibabaseakingtaskreportAPIRequest) GetReportDetail() []TaskDetailReportDto {
	return r._reportDetail
}

// SetTaskType is TaskType Setter
// 任务类型(title/image)
func (r *AlibabaseakingtaskreportAPIRequest) SetTaskType(_taskType string) error {
	r._taskType = _taskType
	r.Set("task_type", _taskType)
	return nil
}

// GetTaskType TaskType Getter
func (r AlibabaseakingtaskreportAPIRequest) GetTaskType() string {
	return r._taskType
}

// SetToken is Token Setter
// 用户token
func (r *AlibabaseakingtaskreportAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaseakingtaskreportAPIRequest) GetToken() string {
	return r._token
}

// SetTokenFrom is TokenFrom Setter
// token来源站点
func (r *AlibabaseakingtaskreportAPIRequest) SetTokenFrom(_tokenFrom string) error {
	r._tokenFrom = _tokenFrom
	r.Set("token_from", _tokenFrom)
	return nil
}

// GetTokenFrom TokenFrom Getter
func (r AlibabaseakingtaskreportAPIRequest) GetTokenFrom() string {
	return r._tokenFrom
}
