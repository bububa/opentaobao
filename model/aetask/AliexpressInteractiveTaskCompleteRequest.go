package aetask

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
任务完成接口 APIRequest
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

func NewAliexpressInteractiveTaskCompleteRequest() *AliexpressInteractiveTaskCompleteRequest{
    return &AliexpressInteractiveTaskCompleteRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressInteractiveTaskCompleteRequest) GetApiMethodName() string {
    return "aliexpress.interactive.task.complete"
}

func (r AliexpressInteractiveTaskCompleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressInteractiveTaskCompleteRequest) SetTaskInstanceId(taskInstanceId int64) error {
    r.taskInstanceId = taskInstanceId
    r.Set("task_instance_id", taskInstanceId)
    return nil
}

func (r AliexpressInteractiveTaskCompleteRequest) GetTaskInstanceId() int64 {
    return r.taskInstanceId
}

func (r *AliexpressInteractiveTaskCompleteRequest) SetProjectAppKey(projectAppKey string) error {
    r.projectAppKey = projectAppKey
    r.Set("project_app_key", projectAppKey)
    return nil
}

func (r AliexpressInteractiveTaskCompleteRequest) GetProjectAppKey() string {
    return r.projectAppKey
}

