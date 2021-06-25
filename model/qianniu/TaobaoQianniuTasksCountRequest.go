package qianniu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
任务查询条数接口 APIRequest
taobao.qianniu.tasks.count

任务查询条数接口
*/
type TaobaoQianniuTasksCountRequest struct {
    model.Params

    // 业务类型
    bizType   string 

    // 子任务类型
    subBizType   string 

    // 业务ID列表，逗号分隔
    bizIds   string 

    // 任务的ID列表，用逗号分隔
    taskIds   string 

    // 任务发起者用户数字ID
    senderUid   int64 

    // 任务执行者用户数字ID
    receiverUid   int64 

    // 逗号分隔的任务状态：0-未执行，1-执行中，2-执行完成，3-超时，4-取消，5-忽略
    status   string 

    // 逗号分隔的子任务状态，由业务方自定义
    subStatus   string 

    // 0-不需要提醒，未设提醒时间 1-设置过提醒时间，需要提醒
    remindFlag   int64 

    // 任务元id，多个以逗号分隔
    metadataIds   string 

    // 与业务相关的买家nick
    bizNick   string 

    // 按时间段搜索时的开始日期，格式如2014-01-01，不填则不限
    startDate   string 

    // 按时间段搜索的结束日期。不填则不限。格式为2014-01-01
    endDate   string 

    // 优先级
    priority   int64 

    // 需要排除的任务类型
    excludeBizType   string 

    // 关键词搜索。只对任务内容进行模糊匹配，以及bizid和biznick进行精准匹配
    keyWord   string 

}

func NewTaobaoQianniuTasksCountRequest() *TaobaoQianniuTasksCountRequest{
    return &TaobaoQianniuTasksCountRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQianniuTasksCountRequest) GetApiMethodName() string {
    return "taobao.qianniu.tasks.count"
}

func (r TaobaoQianniuTasksCountRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQianniuTasksCountRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

func (r TaobaoQianniuTasksCountRequest) GetBizType() string {
    return r.bizType
}

func (r *TaobaoQianniuTasksCountRequest) SetSubBizType(subBizType string) error {
    r.subBizType = subBizType
    r.Set("sub_biz_type", subBizType)
    return nil
}

func (r TaobaoQianniuTasksCountRequest) GetSubBizType() string {
    return r.subBizType
}

func (r *TaobaoQianniuTasksCountRequest) SetBizIds(bizIds string) error {
    r.bizIds = bizIds
    r.Set("biz_ids", bizIds)
    return nil
}

func (r TaobaoQianniuTasksCountRequest) GetBizIds() string {
    return r.bizIds
}

func (r *TaobaoQianniuTasksCountRequest) SetTaskIds(taskIds string) error {
    r.taskIds = taskIds
    r.Set("task_ids", taskIds)
    return nil
}

func (r TaobaoQianniuTasksCountRequest) GetTaskIds() string {
    return r.taskIds
}

func (r *TaobaoQianniuTasksCountRequest) SetSenderUid(senderUid int64) error {
    r.senderUid = senderUid
    r.Set("sender_uid", senderUid)
    return nil
}

func (r TaobaoQianniuTasksCountRequest) GetSenderUid() int64 {
    return r.senderUid
}

func (r *TaobaoQianniuTasksCountRequest) SetReceiverUid(receiverUid int64) error {
    r.receiverUid = receiverUid
    r.Set("receiver_uid", receiverUid)
    return nil
}

func (r TaobaoQianniuTasksCountRequest) GetReceiverUid() int64 {
    return r.receiverUid
}

func (r *TaobaoQianniuTasksCountRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoQianniuTasksCountRequest) GetStatus() string {
    return r.status
}

func (r *TaobaoQianniuTasksCountRequest) SetSubStatus(subStatus string) error {
    r.subStatus = subStatus
    r.Set("sub_status", subStatus)
    return nil
}

func (r TaobaoQianniuTasksCountRequest) GetSubStatus() string {
    return r.subStatus
}

func (r *TaobaoQianniuTasksCountRequest) SetRemindFlag(remindFlag int64) error {
    r.remindFlag = remindFlag
    r.Set("remind_flag", remindFlag)
    return nil
}

func (r TaobaoQianniuTasksCountRequest) GetRemindFlag() int64 {
    return r.remindFlag
}

func (r *TaobaoQianniuTasksCountRequest) SetMetadataIds(metadataIds string) error {
    r.metadataIds = metadataIds
    r.Set("metadata_ids", metadataIds)
    return nil
}

func (r TaobaoQianniuTasksCountRequest) GetMetadataIds() string {
    return r.metadataIds
}

func (r *TaobaoQianniuTasksCountRequest) SetBizNick(bizNick string) error {
    r.bizNick = bizNick
    r.Set("biz_nick", bizNick)
    return nil
}

func (r TaobaoQianniuTasksCountRequest) GetBizNick() string {
    return r.bizNick
}

func (r *TaobaoQianniuTasksCountRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

func (r TaobaoQianniuTasksCountRequest) GetStartDate() string {
    return r.startDate
}

func (r *TaobaoQianniuTasksCountRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r TaobaoQianniuTasksCountRequest) GetEndDate() string {
    return r.endDate
}

func (r *TaobaoQianniuTasksCountRequest) SetPriority(priority int64) error {
    r.priority = priority
    r.Set("priority", priority)
    return nil
}

func (r TaobaoQianniuTasksCountRequest) GetPriority() int64 {
    return r.priority
}

func (r *TaobaoQianniuTasksCountRequest) SetExcludeBizType(excludeBizType string) error {
    r.excludeBizType = excludeBizType
    r.Set("exclude_biz_type", excludeBizType)
    return nil
}

func (r TaobaoQianniuTasksCountRequest) GetExcludeBizType() string {
    return r.excludeBizType
}

func (r *TaobaoQianniuTasksCountRequest) SetKeyWord(keyWord string) error {
    r.keyWord = keyWord
    r.Set("key_word", keyWord)
    return nil
}

func (r TaobaoQianniuTasksCountRequest) GetKeyWord() string {
    return r.keyWord
}

