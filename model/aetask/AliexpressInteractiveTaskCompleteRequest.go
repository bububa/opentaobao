package aetask

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
任务完成接口 API请求
aliexpress.interactive.task.complete

用户完成任务
*/
type AliexpressInteractiveTaskCompleteAPIRequest struct {
    model.Params
    // 任务实例id
    _taskInstanceId   int64
    // appkey
    _projectAppKey   string
}

// 初始化AliexpressInteractiveTaskCompleteAPIRequest对象
func NewAliexpressInteractiveTaskCompleteRequest() *AliexpressInteractiveTaskCompleteAPIRequest{
    return &AliexpressInteractiveTaskCompleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressInteractiveTaskCompleteAPIRequest) GetApiMethodName() string {
    return "aliexpress.interactive.task.complete"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressInteractiveTaskCompleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TaskInstanceId Setter
// 任务实例id
func (r *AliexpressInteractiveTaskCompleteAPIRequest) SetTaskInstanceId(_taskInstanceId int64) error {
    r._taskInstanceId = _taskInstanceId
    r.Set("task_instance_id", _taskInstanceId)
    return nil
}

// TaskInstanceId Getter
func (r AliexpressInteractiveTaskCompleteAPIRequest) GetTaskInstanceId() int64 {
    return r._taskInstanceId
}
// ProjectAppKey Setter
// appkey
func (r *AliexpressInteractiveTaskCompleteAPIRequest) SetProjectAppKey(_projectAppKey string) error {
    r._projectAppKey = _projectAppKey
    r.Set("project_app_key", _projectAppKey)
    return nil
}

// ProjectAppKey Getter
func (r AliexpressInteractiveTaskCompleteAPIRequest) GetProjectAppKey() string {
    return r._projectAppKey
}
