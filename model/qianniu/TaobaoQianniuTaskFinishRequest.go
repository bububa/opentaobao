package qianniu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
完成轻任务 API请求
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

// 初始化TaobaoQianniuTaskFinishRequest对象
func NewTaobaoQianniuTaskFinishRequest() *TaobaoQianniuTaskFinishRequest{
    return &TaobaoQianniuTaskFinishRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQianniuTaskFinishRequest) GetApiMethodName() string {
    return "taobao.qianniu.task.finish"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQianniuTaskFinishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TaskId Setter
// 任务ID
func (r *TaobaoQianniuTaskFinishRequest) SetTaskId(taskId int64) error {
    r.taskId = taskId
    r.Set("task_id", taskId)
    return nil
}

// TaskId Getter
func (r TaobaoQianniuTaskFinishRequest) GetTaskId() int64 {
    return r.taskId
}
// Memo Setter
// 任务备注
func (r *TaobaoQianniuTaskFinishRequest) SetMemo(memo string) error {
    r.memo = memo
    r.Set("memo", memo)
    return nil
}

// Memo Getter
func (r TaobaoQianniuTaskFinishRequest) GetMemo() string {
    return r.memo
}
