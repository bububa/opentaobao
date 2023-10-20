package aetask

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressInteractiveTaskCompleteAPIRequest 任务完成接口 API请求
// aliexpress.interactive.task.complete
//
// 用户完成任务
type AliexpressInteractiveTaskCompleteAPIRequest struct {
	model.Params
	// appkey
	_projectAppKey string
	// 任务实例id
	_taskInstanceId int64
}

// NewAliexpressInteractiveTaskCompleteRequest 初始化AliexpressInteractiveTaskCompleteAPIRequest对象
func NewAliexpressInteractiveTaskCompleteRequest() *AliexpressInteractiveTaskCompleteAPIRequest {
	return &AliexpressInteractiveTaskCompleteAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressInteractiveTaskCompleteAPIRequest) Reset() {
	r._projectAppKey = ""
	r._taskInstanceId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressInteractiveTaskCompleteAPIRequest) GetApiMethodName() string {
	return "aliexpress.interactive.task.complete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressInteractiveTaskCompleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressInteractiveTaskCompleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProjectAppKey is ProjectAppKey Setter
// appkey
func (r *AliexpressInteractiveTaskCompleteAPIRequest) SetProjectAppKey(_projectAppKey string) error {
	r._projectAppKey = _projectAppKey
	r.Set("project_app_key", _projectAppKey)
	return nil
}

// GetProjectAppKey ProjectAppKey Getter
func (r AliexpressInteractiveTaskCompleteAPIRequest) GetProjectAppKey() string {
	return r._projectAppKey
}

// SetTaskInstanceId is TaskInstanceId Setter
// 任务实例id
func (r *AliexpressInteractiveTaskCompleteAPIRequest) SetTaskInstanceId(_taskInstanceId int64) error {
	r._taskInstanceId = _taskInstanceId
	r.Set("task_instance_id", _taskInstanceId)
	return nil
}

// GetTaskInstanceId TaskInstanceId Getter
func (r AliexpressInteractiveTaskCompleteAPIRequest) GetTaskInstanceId() int64 {
	return r._taskInstanceId
}

var poolAliexpressInteractiveTaskCompleteAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressInteractiveTaskCompleteRequest()
	},
}

// GetAliexpressInteractiveTaskCompleteRequest 从 sync.Pool 获取 AliexpressInteractiveTaskCompleteAPIRequest
func GetAliexpressInteractiveTaskCompleteAPIRequest() *AliexpressInteractiveTaskCompleteAPIRequest {
	return poolAliexpressInteractiveTaskCompleteAPIRequest.Get().(*AliexpressInteractiveTaskCompleteAPIRequest)
}

// ReleaseAliexpressInteractiveTaskCompleteAPIRequest 将 AliexpressInteractiveTaskCompleteAPIRequest 放入 sync.Pool
func ReleaseAliexpressInteractiveTaskCompleteAPIRequest(v *AliexpressInteractiveTaskCompleteAPIRequest) {
	v.Reset()
	poolAliexpressInteractiveTaskCompleteAPIRequest.Put(v)
}
