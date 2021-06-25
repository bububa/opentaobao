package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物流管控作业状态回传 APIRequest
alibaba.wdk.fulfill.batch.on.task.status.changed

物流管控作业状态回传
*/
type AlibabaWdkFulfillBatchOnTaskStatusChangedRequest struct {
    model.Params

    // 作业状态回传对象
    taskStatus   *TaskStatus 

}

func NewAlibabaWdkFulfillBatchOnTaskStatusChangedRequest() *AlibabaWdkFulfillBatchOnTaskStatusChangedRequest{
    return &AlibabaWdkFulfillBatchOnTaskStatusChangedRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkFulfillBatchOnTaskStatusChangedRequest) GetApiMethodName() string {
    return "alibaba.wdk.fulfill.batch.on.task.status.changed"
}

func (r AlibabaWdkFulfillBatchOnTaskStatusChangedRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkFulfillBatchOnTaskStatusChangedRequest) SetTaskStatus(taskStatus *TaskStatus) error {
    r.taskStatus = taskStatus
    r.Set("task_status", taskStatus)
    return nil
}

func (r AlibabaWdkFulfillBatchOnTaskStatusChangedRequest) GetTaskStatus() *TaskStatus {
    return r.taskStatus
}

