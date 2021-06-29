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
type AliexpressInteractiveTaskCompleteRequest struct {
    model.Params
    // 任务实例id
    _taskInstanceId   int64
    // appkey
    _projectAppKey   string
}

// 初始化AliexpressInteractiveTaskCompleteRequest对象
func NewAliexpressInteractiveTaskCompleteRequest() *AliexpressInteractiveTaskCompleteRequest{
    return &AliexpressInteractiveTaskCompleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressInteractiveTaskCompleteRequest) GetApiMethodName() string {
    return "aliexpress.interactive.task.complete"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressInteractiveTaskCompleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TaskInstanceId Setter
// 任务实例id
func (r *AliexpressInteractiveTaskCompleteRequest) SetTaskInstanceId(_taskInstanceId int64) error {
    r._taskInstanceId = _taskInstanceId
    r.Set("task_instance_id", _taskInstanceId)
    return nil
}

// TaskInstanceId Getter
func (r AliexpressInteractiveTaskCompleteRequest) GetTaskInstanceId() int64 {
    return r._taskInstanceId
}
// ProjectAppKey Setter
// appkey
func (r *AliexpressInteractiveTaskCompleteRequest) SetProjectAppKey(_projectAppKey string) error {
    r._projectAppKey = _projectAppKey
    r.Set("project_app_key", _projectAppKey)
    return nil
}

// ProjectAppKey Getter
func (r AliexpressInteractiveTaskCompleteRequest) GetProjectAppKey() string {
    return r._projectAppKey
}
