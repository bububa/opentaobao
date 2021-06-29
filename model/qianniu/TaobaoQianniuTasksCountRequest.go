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
    _bizType   string
    // 子任务类型
    _subBizType   string
    // 业务ID列表，逗号分隔
    _bizIds   string
    // 任务的ID列表，用逗号分隔
    _taskIds   string
    // 任务发起者用户数字ID
    _senderUid   int64
    // 任务执行者用户数字ID
    _receiverUid   int64
    // 逗号分隔的任务状态：0-未执行，1-执行中，2-执行完成，3-超时，4-取消，5-忽略
    _status   string
    // 逗号分隔的子任务状态，由业务方自定义
    _subStatus   string
    // 0-不需要提醒，未设提醒时间 1-设置过提醒时间，需要提醒
    _remindFlag   int64
    // 任务元id，多个以逗号分隔
    _metadataIds   string
    // 与业务相关的买家nick
    _bizNick   string
    // 按时间段搜索时的开始日期，格式如2014-01-01，不填则不限
    _startDate   string
    // 按时间段搜索的结束日期。不填则不限。格式为2014-01-01
    _endDate   string
    // 优先级
    _priority   int64
    // 需要排除的任务类型
    _excludeBizType   string
    // 关键词搜索。只对任务内容进行模糊匹配，以及bizid和biznick进行精准匹配
    _keyWord   string
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
func (r *TaobaoQianniuTasksCountRequest) SetBizType(_bizType string) error {
    r._bizType = _bizType
    r.Set("biz_type", _bizType)
    return nil
}

// BizType Getter
func (r TaobaoQianniuTasksCountRequest) GetBizType() string {
    return r._bizType
}
// SubBizType Setter
// 子任务类型
func (r *TaobaoQianniuTasksCountRequest) SetSubBizType(_subBizType string) error {
    r._subBizType = _subBizType
    r.Set("sub_biz_type", _subBizType)
    return nil
}

// SubBizType Getter
func (r TaobaoQianniuTasksCountRequest) GetSubBizType() string {
    return r._subBizType
}
// BizIds Setter
// 业务ID列表，逗号分隔
func (r *TaobaoQianniuTasksCountRequest) SetBizIds(_bizIds string) error {
    r._bizIds = _bizIds
    r.Set("biz_ids", _bizIds)
    return nil
}

// BizIds Getter
func (r TaobaoQianniuTasksCountRequest) GetBizIds() string {
    return r._bizIds
}
// TaskIds Setter
// 任务的ID列表，用逗号分隔
func (r *TaobaoQianniuTasksCountRequest) SetTaskIds(_taskIds string) error {
    r._taskIds = _taskIds
    r.Set("task_ids", _taskIds)
    return nil
}

// TaskIds Getter
func (r TaobaoQianniuTasksCountRequest) GetTaskIds() string {
    return r._taskIds
}
// SenderUid Setter
// 任务发起者用户数字ID
func (r *TaobaoQianniuTasksCountRequest) SetSenderUid(_senderUid int64) error {
    r._senderUid = _senderUid
    r.Set("sender_uid", _senderUid)
    return nil
}

// SenderUid Getter
func (r TaobaoQianniuTasksCountRequest) GetSenderUid() int64 {
    return r._senderUid
}
// ReceiverUid Setter
// 任务执行者用户数字ID
func (r *TaobaoQianniuTasksCountRequest) SetReceiverUid(_receiverUid int64) error {
    r._receiverUid = _receiverUid
    r.Set("receiver_uid", _receiverUid)
    return nil
}

// ReceiverUid Getter
func (r TaobaoQianniuTasksCountRequest) GetReceiverUid() int64 {
    return r._receiverUid
}
// Status Setter
// 逗号分隔的任务状态：0-未执行，1-执行中，2-执行完成，3-超时，4-取消，5-忽略
func (r *TaobaoQianniuTasksCountRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoQianniuTasksCountRequest) GetStatus() string {
    return r._status
}
// SubStatus Setter
// 逗号分隔的子任务状态，由业务方自定义
func (r *TaobaoQianniuTasksCountRequest) SetSubStatus(_subStatus string) error {
    r._subStatus = _subStatus
    r.Set("sub_status", _subStatus)
    return nil
}

// SubStatus Getter
func (r TaobaoQianniuTasksCountRequest) GetSubStatus() string {
    return r._subStatus
}
// RemindFlag Setter
// 0-不需要提醒，未设提醒时间 1-设置过提醒时间，需要提醒
func (r *TaobaoQianniuTasksCountRequest) SetRemindFlag(_remindFlag int64) error {
    r._remindFlag = _remindFlag
    r.Set("remind_flag", _remindFlag)
    return nil
}

// RemindFlag Getter
func (r TaobaoQianniuTasksCountRequest) GetRemindFlag() int64 {
    return r._remindFlag
}
// MetadataIds Setter
// 任务元id，多个以逗号分隔
func (r *TaobaoQianniuTasksCountRequest) SetMetadataIds(_metadataIds string) error {
    r._metadataIds = _metadataIds
    r.Set("metadata_ids", _metadataIds)
    return nil
}

// MetadataIds Getter
func (r TaobaoQianniuTasksCountRequest) GetMetadataIds() string {
    return r._metadataIds
}
// BizNick Setter
// 与业务相关的买家nick
func (r *TaobaoQianniuTasksCountRequest) SetBizNick(_bizNick string) error {
    r._bizNick = _bizNick
    r.Set("biz_nick", _bizNick)
    return nil
}

// BizNick Getter
func (r TaobaoQianniuTasksCountRequest) GetBizNick() string {
    return r._bizNick
}
// StartDate Setter
// 按时间段搜索时的开始日期，格式如2014-01-01，不填则不限
func (r *TaobaoQianniuTasksCountRequest) SetStartDate(_startDate string) error {
    r._startDate = _startDate
    r.Set("start_date", _startDate)
    return nil
}

// StartDate Getter
func (r TaobaoQianniuTasksCountRequest) GetStartDate() string {
    return r._startDate
}
// EndDate Setter
// 按时间段搜索的结束日期。不填则不限。格式为2014-01-01
func (r *TaobaoQianniuTasksCountRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r TaobaoQianniuTasksCountRequest) GetEndDate() string {
    return r._endDate
}
// Priority Setter
// 优先级
func (r *TaobaoQianniuTasksCountRequest) SetPriority(_priority int64) error {
    r._priority = _priority
    r.Set("priority", _priority)
    return nil
}

// Priority Getter
func (r TaobaoQianniuTasksCountRequest) GetPriority() int64 {
    return r._priority
}
// ExcludeBizType Setter
// 需要排除的任务类型
func (r *TaobaoQianniuTasksCountRequest) SetExcludeBizType(_excludeBizType string) error {
    r._excludeBizType = _excludeBizType
    r.Set("exclude_biz_type", _excludeBizType)
    return nil
}

// ExcludeBizType Getter
func (r TaobaoQianniuTasksCountRequest) GetExcludeBizType() string {
    return r._excludeBizType
}
// KeyWord Setter
// 关键词搜索。只对任务内容进行模糊匹配，以及bizid和biznick进行精准匹配
func (r *TaobaoQianniuTasksCountRequest) SetKeyWord(_keyWord string) error {
    r._keyWord = _keyWord
    r.Set("key_word", _keyWord)
    return nil
}

// KeyWord Getter
func (r TaobaoQianniuTasksCountRequest) GetKeyWord() string {
    return r._keyWord
}
