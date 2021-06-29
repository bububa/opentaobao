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
    taskInstanceId   int64
    // appkey
    projectAppKey   string
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
func (r *AliexpressInteractiveTaskCompleteRequest) SetTaskInstanceId(taskInstanceId int64) error {
    r.taskInstanceId = taskInstanceId
    r.Set("task_instance_id", taskInstanceId)
    return nil
}

// TaskInstanceId Getter
func (r AliexpressInteractiveTaskCompleteRequest) GetTaskInstanceId() int64 {
    return r.taskInstanceId
}
// ProjectAppKey Setter
// appkey
func (r *AliexpressInteractiveTaskCompleteRequest) SetProjectAppKey(projectAppKey string) error {
    r.projectAppKey = projectAppKey
    r.Set("project_app_key", projectAppKey)
    return nil
}

// ProjectAppKey Getter
func (r AliexpressInteractiveTaskCompleteRequest) GetProjectAppKey() string {
    return r.projectAppKey
}
