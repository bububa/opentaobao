package qianniu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
完成轻任务 APIRequest
taobao.qianniu.task.finish

由任务执行者调用
*/
type TaobaoQianniuTaskFinishRequest struct {
    model.Params

    // 任务ID
    taskId   int64 

    // 任务备注
    memo   string 

}

func NewTaobaoQianniuTaskFinishRequest() *TaobaoQianniuTaskFinishRequest{
    return &TaobaoQianniuTaskFinishRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQianniuTaskFinishRequest) GetApiMethodName() string {
    return "taobao.qianniu.task.finish"
}

func (r TaobaoQianniuTaskFinishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQianniuTaskFinishRequest) SetTaskId(taskId int64) error {
    r.taskId = taskId
    r.Set("task_id", taskId)
    return nil
}

func (r TaobaoQianniuTaskFinishRequest) GetTaskId() int64 {
    return r.taskId
}

func (r *TaobaoQianniuTaskFinishRequest) SetMemo(memo string) error {
    r.memo = memo
    r.Set("memo", memo)
    return nil
}

func (r TaobaoQianniuTaskFinishRequest) GetMemo() string {
    return r.memo
}

