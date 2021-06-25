package qianniu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新轻任务 APIRequest
taobao.qianniu.task.update

由任务执行者调用，sub_status，tag和memo至少提供一个
*/
type TaobaoQianniuTaskUpdateRequest struct {
    model.Params

    // 任务ID
    taskId   int64 

    // 子任务状态，由业务方自定义
    subStatus   string 

    // 任务标签
    tag   string 

    // 任务备注。当memo_mode为1时，memo将采用追加方式。
    memo   string 

    // 状态值，多个以逗号分隔
    status   string 

    // 提醒时间，时间的毫秒数
    remindTime   int64 

    // 应用自定义参数
    bizParam   string 

    // 0为不提醒，1为全部提醒，2为PC提醒，3为移动提醒，4为已提醒，5为已忽略。
    remindFlag   int64 

    // 表示memo字段的更新策略。如需采用追加方式的，请将此字段设置为1。
    memoMode   int64 

    // 默认填0，数字越大优化级越高。当前常用0和1.
    priority   int64 

    // 0表示没有删除，1表示删除
    isDeleted   int64 

}

func NewTaobaoQianniuTaskUpdateRequest() *TaobaoQianniuTaskUpdateRequest{
    return &TaobaoQianniuTaskUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQianniuTaskUpdateRequest) GetApiMethodName() string {
    return "taobao.qianniu.task.update"
}

func (r TaobaoQianniuTaskUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQianniuTaskUpdateRequest) SetTaskId(taskId int64) error {
    r.taskId = taskId
    r.Set("task_id", taskId)
    return nil
}

func (r TaobaoQianniuTaskUpdateRequest) GetTaskId() int64 {
    return r.taskId
}

func (r *TaobaoQianniuTaskUpdateRequest) SetSubStatus(subStatus string) error {
    r.subStatus = subStatus
    r.Set("sub_status", subStatus)
    return nil
}

func (r TaobaoQianniuTaskUpdateRequest) GetSubStatus() string {
    return r.subStatus
}

func (r *TaobaoQianniuTaskUpdateRequest) SetTag(tag string) error {
    r.tag = tag
    r.Set("tag", tag)
    return nil
}

func (r TaobaoQianniuTaskUpdateRequest) GetTag() string {
    return r.tag
}

func (r *TaobaoQianniuTaskUpdateRequest) SetMemo(memo string) error {
    r.memo = memo
    r.Set("memo", memo)
    return nil
}

func (r TaobaoQianniuTaskUpdateRequest) GetMemo() string {
    return r.memo
}

func (r *TaobaoQianniuTaskUpdateRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoQianniuTaskUpdateRequest) GetStatus() string {
    return r.status
}

func (r *TaobaoQianniuTaskUpdateRequest) SetRemindTime(remindTime int64) error {
    r.remindTime = remindTime
    r.Set("remind_time", remindTime)
    return nil
}

func (r TaobaoQianniuTaskUpdateRequest) GetRemindTime() int64 {
    return r.remindTime
}

func (r *TaobaoQianniuTaskUpdateRequest) SetBizParam(bizParam string) error {
    r.bizParam = bizParam
    r.Set("biz_param", bizParam)
    return nil
}

func (r TaobaoQianniuTaskUpdateRequest) GetBizParam() string {
    return r.bizParam
}

func (r *TaobaoQianniuTaskUpdateRequest) SetRemindFlag(remindFlag int64) error {
    r.remindFlag = remindFlag
    r.Set("remind_flag", remindFlag)
    return nil
}

func (r TaobaoQianniuTaskUpdateRequest) GetRemindFlag() int64 {
    return r.remindFlag
}

func (r *TaobaoQianniuTaskUpdateRequest) SetMemoMode(memoMode int64) error {
    r.memoMode = memoMode
    r.Set("memo_mode", memoMode)
    return nil
}

func (r TaobaoQianniuTaskUpdateRequest) GetMemoMode() int64 {
    return r.memoMode
}

func (r *TaobaoQianniuTaskUpdateRequest) SetPriority(priority int64) error {
    r.priority = priority
    r.Set("priority", priority)
    return nil
}

func (r TaobaoQianniuTaskUpdateRequest) GetPriority() int64 {
    return r.priority
}

func (r *TaobaoQianniuTaskUpdateRequest) SetIsDeleted(isDeleted int64) error {
    r.isDeleted = isDeleted
    r.Set("is_deleted", isDeleted)
    return nil
}

func (r TaobaoQianniuTaskUpdateRequest) GetIsDeleted() int64 {
    return r.isDeleted
}

