package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkfulfillbatchontaskstatuschangedAPIRequest 物流管控作业状态回传 API请求
// alibaba.wdk.fulfill.batch.on.task.status.changed
//
// 物流管控作业状态回传
type AlibabawdkfulfillbatchontaskstatuschangedAPIRequest struct {
	model.Params
	// 作业状态回传对象
	_taskStatus *TaskStatus
}

// NewAlibabawdkfulfillbatchontaskstatuschangedRequest 初始化AlibabawdkfulfillbatchontaskstatuschangedAPIRequest对象
func NewAlibabawdkfulfillbatchontaskstatuschangedRequest() *AlibabawdkfulfillbatchontaskstatuschangedAPIRequest {
	return &AlibabawdkfulfillbatchontaskstatuschangedAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkfulfillbatchontaskstatuschangedAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.fulfill.batch.on.task.status.changed"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkfulfillbatchontaskstatuschangedAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkfulfillbatchontaskstatuschangedAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTaskStatus is TaskStatus Setter
// 作业状态回传对象
func (r *AlibabawdkfulfillbatchontaskstatuschangedAPIRequest) SetTaskStatus(_taskStatus *TaskStatus) error {
	r._taskStatus = _taskStatus
	r.Set("task_status", _taskStatus)
	return nil
}

// GetTaskStatus TaskStatus Getter
func (r AlibabawdkfulfillbatchontaskstatuschangedAPIRequest) GetTaskStatus() *TaskStatus {
	return r._taskStatus
}
