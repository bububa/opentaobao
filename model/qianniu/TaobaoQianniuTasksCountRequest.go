package qianniu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
任务查询条数接口 API请求
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

// 初始化TaobaoQianniuTasksCountRequest对象
func NewTaobaoQianniuTasksCountRequest() *TaobaoQianniuTasksCountRequest{
    return &TaobaoQianniuTasksCountRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQianniuTasksCountRequest) GetApiMethodName() string {
    return "taobao.qianniu.tasks.count"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQianniuTasksCountRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizType Setter
// 业务类型
func (r *TaobaoQianniuTasksCountRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

// BizType Getter
func (r TaobaoQianniuTasksCountRequest) GetBizType() string {
    return r.bizType
}
// SubBizType Setter
// 子任务类型
func (r *TaobaoQianniuTasksCountRequest) SetSubBizType(subBizType string) error {
    r.subBizType = subBizType
    r.Set("sub_biz_type", subBizType)
    return nil
}

// SubBizType Getter
func (r TaobaoQianniuTasksCountRequest) GetSubBizType() string {
    return r.subBizType
}
// BizIds Setter
// 业务ID列表，逗号分隔
func (r *TaobaoQianniuTasksCountRequest) SetBizIds(bizIds string) error {
    r.bizIds = bizIds
    r.Set("biz_ids", bizIds)
    return nil
}

// BizIds Getter
func (r TaobaoQianniuTasksCountRequest) GetBizIds() string {
    return r.bizIds
}
// TaskIds Setter
// 任务的ID列表，用逗号分隔
func (r *TaobaoQianniuTasksCountRequest) SetTaskIds(taskIds string) error {
    r.taskIds = taskIds
    r.Set("task_ids", taskIds)
    return nil
}

// TaskIds Getter
func (r TaobaoQianniuTasksCountRequest) GetTaskIds() string {
    return r.taskIds
}
// SenderUid Setter
// 任务发起者用户数字ID
func (r *TaobaoQianniuTasksCountRequest) SetSenderUid(senderUid int64) error {
    r.senderUid = senderUid
    r.Set("sender_uid", senderUid)
    return nil
}

// SenderUid Getter
func (r TaobaoQianniuTasksCountRequest) GetSenderUid() int64 {
    return r.senderUid
}
// ReceiverUid Setter
// 任务执行者用户数字ID
func (r *TaobaoQianniuTasksCountRequest) SetReceiverUid(receiverUid int64) error {
    r.receiverUid = receiverUid
    r.Set("receiver_uid", receiverUid)
    return nil
}

// ReceiverUid Getter
func (r TaobaoQianniuTasksCountRequest) GetReceiverUid() int64 {
    return r.receiverUid
}
// Status Setter
// 逗号分隔的任务状态：0-未执行，1-执行中，2-执行完成，3-超时，4-取消，5-忽略
func (r *TaobaoQianniuTasksCountRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoQianniuTasksCountRequest) GetStatus() string {
    return r.status
}
// SubStatus Setter
// 逗号分隔的子任务状态，由业务方自定义
func (r *TaobaoQianniuTasksCountRequest) SetSubStatus(subStatus string) error {
    r.subStatus = subStatus
    r.Set("sub_status", subStatus)
    return nil
}

// SubStatus Getter
func (r TaobaoQianniuTasksCountRequest) GetSubStatus() string {
    return r.subStatus
}
// RemindFlag Setter
// 0-不需要提醒，未设提醒时间 1-设置过提醒时间，需要提醒
func (r *TaobaoQianniuTasksCountRequest) SetRemindFlag(remindFlag int64) error {
    r.remindFlag = remindFlag
    r.Set("remind_flag", remindFlag)
    return nil
}

// RemindFlag Getter
func (r TaobaoQianniuTasksCountRequest) GetRemindFlag() int64 {
    return r.remindFlag
}
// MetadataIds Setter
// 任务元id，多个以逗号分隔
func (r *TaobaoQianniuTasksCountRequest) SetMetadataIds(metadataIds string) error {
    r.metadataIds = metadataIds
    r.Set("metadata_ids", metadataIds)
    return nil
}

// MetadataIds Getter
func (r TaobaoQianniuTasksCountRequest) GetMetadataIds() string {
    return r.metadataIds
}
// BizNick Setter
// 与业务相关的买家nick
func (r *TaobaoQianniuTasksCountRequest) SetBizNick(bizNick string) error {
    r.bizNick = bizNick
    r.Set("biz_nick", bizNick)
    return nil
}

// BizNick Getter
func (r TaobaoQianniuTasksCountRequest) GetBizNick() string {
    return r.bizNick
}
// StartDate Setter
// 按时间段搜索时的开始日期，格式如2014-01-01，不填则不限
func (r *TaobaoQianniuTasksCountRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

// StartDate Getter
func (r TaobaoQianniuTasksCountRequest) GetStartDate() string {
    return r.startDate
}
// EndDate Setter
// 按时间段搜索的结束日期。不填则不限。格式为2014-01-01
func (r *TaobaoQianniuTasksCountRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r TaobaoQianniuTasksCountRequest) GetEndDate() string {
    return r.endDate
}
// Priority Setter
// 优先级
func (r *TaobaoQianniuTasksCountRequest) SetPriority(priority int64) error {
    r.priority = priority
    r.Set("priority", priority)
    return nil
}

// Priority Getter
func (r TaobaoQianniuTasksCountRequest) GetPriority() int64 {
    return r.priority
}
// ExcludeBizType Setter
// 需要排除的任务类型
func (r *TaobaoQianniuTasksCountRequest) SetExcludeBizType(excludeBizType string) error {
    r.excludeBizType = excludeBizType
    r.Set("exclude_biz_type", excludeBizType)
    return nil
}

// ExcludeBizType Getter
func (r TaobaoQianniuTasksCountRequest) GetExcludeBizType() string {
    return r.excludeBizType
}
// KeyWord Setter
// 关键词搜索。只对任务内容进行模糊匹配，以及bizid和biznick进行精准匹配
func (r *TaobaoQianniuTasksCountRequest) SetKeyWord(keyWord string) error {
    r.keyWord = keyWord
    r.Set("key_word", keyWord)
    return nil
}

// KeyWord Getter
func (r TaobaoQianniuTasksCountRequest) GetKeyWord() string {
    return r.keyWord
}
