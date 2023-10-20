package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkFulfillBatchOnTaskStatusChangedAPIRequest) Reset() {
	r._taskStatus = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkFulfillBatchOnTaskStatusChangedAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.fulfill.batch.on.task.status.changed"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkFulfillBatchOnTaskStatusChangedAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkFulfillBatchOnTaskStatusChangedAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTaskStatus is TaskStatus Setter
// 作业状态回传对象
func (r *AlibabaWdkFulfillBatchOnTaskStatusChangedAPIRequest) SetTaskStatus(_taskStatus *TaskStatus) error {
	r._taskStatus = _taskStatus
	r.Set("task_status", _taskStatus)
	return nil
}

// GetTaskStatus TaskStatus Getter
func (r AlibabaWdkFulfillBatchOnTaskStatusChangedAPIRequest) GetTaskStatus() *TaskStatus {
	return r._taskStatus
}

var poolAlibabaWdkFulfillBatchOnTaskStatusChangedAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkFulfillBatchOnTaskStatusChangedRequest()
	},
}

// GetAlibabaWdkFulfillBatchOnTaskStatusChangedRequest 从 sync.Pool 获取 AlibabaWdkFulfillBatchOnTaskStatusChangedAPIRequest
func GetAlibabaWdkFulfillBatchOnTaskStatusChangedAPIRequest() *AlibabaWdkFulfillBatchOnTaskStatusChangedAPIRequest {
	return poolAlibabaWdkFulfillBatchOnTaskStatusChangedAPIRequest.Get().(*AlibabaWdkFulfillBatchOnTaskStatusChangedAPIRequest)
}

// ReleaseAlibabaWdkFulfillBatchOnTaskStatusChangedAPIRequest 将 AlibabaWdkFulfillBatchOnTaskStatusChangedAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkFulfillBatchOnTaskStatusChangedAPIRequest(v *AlibabaWdkFulfillBatchOnTaskStatusChangedAPIRequest) {
	v.Reset()
	poolAlibabaWdkFulfillBatchOnTaskStatusChangedAPIRequest.Put(v)
}
