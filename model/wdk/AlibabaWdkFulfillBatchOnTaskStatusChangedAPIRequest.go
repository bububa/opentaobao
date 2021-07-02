package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkFulfillBatchOnTaskStatusChangedAPIRequest 物流管控作业状态回传 API请求
// alibaba.wdk.fulfill.batch.on.task.status.changed
//
// 物流管控作业状态回传
type AlibabaWdkFulfillBatchOnTaskStatusChangedAPIRequest struct {
	model.Params
	// 作业状态回传对象
	_taskStatus *TaskStatus
}

// NewAlibabaWdkFulfillBatchOnTaskStatusChangedRequest 初始化AlibabaWdkFulfillBatchOnTaskStatusChangedAPIRequest对象
func NewAlibabaWdkFulfillBatchOnTaskStatusChangedRequest() *AlibabaWdkFulfillBatchOnTaskStatusChangedAPIRequest {
	return &AlibabaWdkFulfillBatchOnTaskStatusChangedAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkFulfillBatchOnTaskStatusChangedAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.fulfill.batch.on.task.status.changed"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkFulfillBatchOnTaskStatusChangedAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TaskStatus Setter
// 作业状态回传对象
func (r *AlibabaWdkFulfillBatchOnTaskStatusChangedAPIRequest) SetTaskStatus(_taskStatus *TaskStatus) error {
	r._taskStatus = _taskStatus
	r.Set("task_status", _taskStatus)
	return nil
}

// Get TaskStatus Getter
func (r AlibabaWdkFulfillBatchOnTaskStatusChangedAPIRequest) GetTaskStatus() *TaskStatus {
	return r._taskStatus
}
