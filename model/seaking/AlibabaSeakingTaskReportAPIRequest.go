package seaking

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSeakingTaskReportAPIRequest) Reset() {
	r._reportDetail = r._reportDetail[:0]
	r._taskType = ""
	r._token = ""
	r._tokenFrom = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSeakingTaskReportAPIRequest) GetApiMethodName() string {
	return "alibaba.seaking.task.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSeakingTaskReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSeakingTaskReportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReportDetail is ReportDetail Setter
// 上报数据详情
func (r *AlibabaSeakingTaskReportAPIRequest) SetReportDetail(_reportDetail []TaskDetailReportDto) error {
	r._reportDetail = _reportDetail
	r.Set("report_detail", _reportDetail)
	return nil
}

// GetReportDetail ReportDetail Getter
func (r AlibabaSeakingTaskReportAPIRequest) GetReportDetail() []TaskDetailReportDto {
	return r._reportDetail
}

// SetTaskType is TaskType Setter
// 任务类型(title/image)
func (r *AlibabaSeakingTaskReportAPIRequest) SetTaskType(_taskType string) error {
	r._taskType = _taskType
	r.Set("task_type", _taskType)
	return nil
}

// GetTaskType TaskType Getter
func (r AlibabaSeakingTaskReportAPIRequest) GetTaskType() string {
	return r._taskType
}

// SetToken is Token Setter
// 用户token
func (r *AlibabaSeakingTaskReportAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaSeakingTaskReportAPIRequest) GetToken() string {
	return r._token
}

// SetTokenFrom is TokenFrom Setter
// token来源站点
func (r *AlibabaSeakingTaskReportAPIRequest) SetTokenFrom(_tokenFrom string) error {
	r._tokenFrom = _tokenFrom
	r.Set("token_from", _tokenFrom)
	return nil
}

// GetTokenFrom TokenFrom Getter
func (r AlibabaSeakingTaskReportAPIRequest) GetTokenFrom() string {
	return r._tokenFrom
}

var poolAlibabaSeakingTaskReportAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSeakingTaskReportRequest()
	},
}

// GetAlibabaSeakingTaskReportRequest 从 sync.Pool 获取 AlibabaSeakingTaskReportAPIRequest
func GetAlibabaSeakingTaskReportAPIRequest() *AlibabaSeakingTaskReportAPIRequest {
	return poolAlibabaSeakingTaskReportAPIRequest.Get().(*AlibabaSeakingTaskReportAPIRequest)
}

// ReleaseAlibabaSeakingTaskReportAPIRequest 将 AlibabaSeakingTaskReportAPIRequest 放入 sync.Pool
func ReleaseAlibabaSeakingTaskReportAPIRequest(v *AlibabaSeakingTaskReportAPIRequest) {
	v.Reset()
	poolAlibabaSeakingTaskReportAPIRequest.Put(v)
}
