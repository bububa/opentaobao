package aetask

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressinteractivetaskcompleteAPIRequest 任务完成接口 API请求
// aliexpress.interactive.task.complete
//
// 用户完成任务
type AliexpressinteractivetaskcompleteAPIRequest struct {
	model.Params
	// appkey
	_projectAppKey string
	// 任务实例id
	_taskInstanceId int64
}

// NewAliexpressinteractivetaskcompleteRequest 初始化AliexpressinteractivetaskcompleteAPIRequest对象
func NewAliexpressinteractivetaskcompleteRequest() *AliexpressinteractivetaskcompleteAPIRequest {
	return &AliexpressinteractivetaskcompleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressinteractivetaskcompleteAPIRequest) GetApiMethodName() string {
	return "aliexpress.interactive.task.complete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressinteractivetaskcompleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressinteractivetaskcompleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProjectAppKey is ProjectAppKey Setter
// appkey
func (r *AliexpressinteractivetaskcompleteAPIRequest) SetProjectAppKey(_projectAppKey string) error {
	r._projectAppKey = _projectAppKey
	r.Set("project_app_key", _projectAppKey)
	return nil
}

// GetProjectAppKey ProjectAppKey Getter
func (r AliexpressinteractivetaskcompleteAPIRequest) GetProjectAppKey() string {
	return r._projectAppKey
}

// SetTaskInstanceId is TaskInstanceId Setter
// 任务实例id
func (r *AliexpressinteractivetaskcompleteAPIRequest) SetTaskInstanceId(_taskInstanceId int64) error {
	r._taskInstanceId = _taskInstanceId
	r.Set("task_instance_id", _taskInstanceId)
	return nil
}

// GetTaskInstanceId TaskInstanceId Getter
func (r AliexpressinteractivetaskcompleteAPIRequest) GetTaskInstanceId() int64 {
	return r._taskInstanceId
}
