package qianniu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取指定的任务 API请求
taobao.qianniu.tasks.get

获取指定的任务，可用的参数组合：<br/>task_ids + need_meta + fields：精确查找<br/>biz_type + sub_biz_type + biz_ids + need_meta + fields：按照业务ID查找<br/>biz_type + sub_biz_type + sender_uid + need_meta + fields：按照发起者查找<br/>biz_type + sub_biz_type + receiver_uid + need_meta + fields：按照执行者查找<br/>biz_type+modify_start_time+modify_end_time+fields:能支持指定修改时间的查询，用于增量查询等
*/
type TaobaoQianniuTasksGetRequest struct {
    model.Params
    // 排序字段，可以为id,gmt_create,gmt_finished,metadata_id等
    _orderBy   string
    // asc为升，desc为降
    _orderType   string
    // 0-不需要提醒，未设提醒时间 1-设置过提醒时间，需要提醒
    _remindFlag   int64
    // 业务相关的对象，当前主要表示买家nick
    _bizNick   string
    // 根据任务创建时间搜索的开始日期（含），不填则不限。例如只查询2014-01-01当天的任务，则将start_date和end_date都设置成2014-01-01
    _startDate   string
    // 根据任务创建时间搜索的结束日期（含），不填则不限。例如只查询2014-01-01当天的任务，则将start_date和end_date都设置成2014-01-01
    _endDate   string
    // 根据任务修改时间搜索的开始时间（含），不填则不限。例如查询“2014-01-01 00:00:10”之后有修改的任务，则将modify_start_time_str设置成“2014-01-01 00:00:10”
    _modifyStartTimeStr   string
    // 根据任务修改时间搜索的结束时间（含），不填则不限。例如查询“2014-01-01 00:00:10”之前有修改的任务，则将modify_end_time_str设置成“2014-01-01 00:00:10”
    _modifyEndTimeStr   string
    // 优先级。即创建时的metadata中的优先级。0为低，1为中，2为高。
    _priority   int64
    // 需要排除的任务类型
    _excludeBizType   string
    // 关键词搜索。只对任务内容进行模糊匹配，以及bizid和biznick进行精准匹配
    _keyWord   string
    // 当前页数，从1开始
    _currentPage   int64
    // 每页条数
    _pageSize   int64
    // 业务类型
    _bizType   string
    // 子任务类型
    _subBizType   string
    // 任务的ID列表，用逗号分隔
    _taskIds   string
    // 业务ID列表，逗号分隔
    _bizIds   string
    // 任务执行者用户数字ID
    _receiverUid   int64
    // 任务发起者用户数字ID
    _senderUid   int64
    // 逗号分隔的任务状态：0-未执行，1-执行中，2-执行完成，3-超时，4-取消，5-忽略
    _status   string
    // 逗号分隔的子任务状态，由业务方自定义
    _subStatus   string
    // 任务元id，多个以逗号分隔
    _metadataIds   string
    // 是否需要meta信息，默认值为false
    _needMeta   bool
    // 逗号分隔的字段列表，各个字段含义： id：任务ID receiver_uid：执行者用户数字ID receiver_nick：执行者用户昵称 status：任务状态：0-未执行，1-执行中，2-执行完成，3-超时，4-取消，5-忽略 sub_status：子任务状态，由业务方自定义 finish_strategy：任务完成策略：1-一个人完成，2-所有人完成 gmt_finished：任务完成时间，格式：时间毫秒数 biz_type：业务类型 sub_biz_type：子业务类型 biz_id：业务ID biz_param：业务参数 biz_entry：业务入口 tag：任务标签 memo：任务备注
    _fields   string
    // 客户端的版本信息
    _clientInfo   string
    // 是否需要删除的任务，默认为false
    _needDeleted   bool
}

// 初始化TaobaoQianniuTasksGetRequest对象
func NewTaobaoQianniuTasksGetRequest() *TaobaoQianniuTasksGetRequest{
    return &TaobaoQianniuTasksGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQianniuTasksGetRequest) GetApiMethodName() string {
    return "taobao.qianniu.tasks.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQianniuTasksGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderBy Setter
// 排序字段，可以为id,gmt_create,gmt_finished,metadata_id等
func (r *TaobaoQianniuTasksGetRequest) SetOrderBy(_orderBy string) error {
    r._orderBy = _orderBy
    r.Set("order_by", _orderBy)
    return nil
}

// OrderBy Getter
func (r TaobaoQianniuTasksGetRequest) GetOrderBy() string {
    return r._orderBy
}
// OrderType Setter
// asc为升，desc为降
func (r *TaobaoQianniuTasksGetRequest) SetOrderType(_orderType string) error {
    r._orderType = _orderType
    r.Set("order_type", _orderType)
    return nil
}

// OrderType Getter
func (r TaobaoQianniuTasksGetRequest) GetOrderType() string {
    return r._orderType
}
// RemindFlag Setter
// 0-不需要提醒，未设提醒时间 1-设置过提醒时间，需要提醒
func (r *TaobaoQianniuTasksGetRequest) SetRemindFlag(_remindFlag int64) error {
    r._remindFlag = _remindFlag
    r.Set("remind_flag", _remindFlag)
    return nil
}

// RemindFlag Getter
func (r TaobaoQianniuTasksGetRequest) GetRemindFlag() int64 {
    return r._remindFlag
}
// BizNick Setter
// 业务相关的对象，当前主要表示买家nick
func (r *TaobaoQianniuTasksGetRequest) SetBizNick(_bizNick string) error {
    r._bizNick = _bizNick
    r.Set("biz_nick", _bizNick)
    return nil
}

// BizNick Getter
func (r TaobaoQianniuTasksGetRequest) GetBizNick() string {
    return r._bizNick
}
// StartDate Setter
// 根据任务创建时间搜索的开始日期（含），不填则不限。例如只查询2014-01-01当天的任务，则将start_date和end_date都设置成2014-01-01
func (r *TaobaoQianniuTasksGetRequest) SetStartDate(_startDate string) error {
    r._startDate = _startDate
    r.Set("start_date", _startDate)
    return nil
}

// StartDate Getter
func (r TaobaoQianniuTasksGetRequest) GetStartDate() string {
    return r._startDate
}
// EndDate Setter
// 根据任务创建时间搜索的结束日期（含），不填则不限。例如只查询2014-01-01当天的任务，则将start_date和end_date都设置成2014-01-01
func (r *TaobaoQianniuTasksGetRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r TaobaoQianniuTasksGetRequest) GetEndDate() string {
    return r._endDate
}
// ModifyStartTimeStr Setter
// 根据任务修改时间搜索的开始时间（含），不填则不限。例如查询“2014-01-01 00:00:10”之后有修改的任务，则将modify_start_time_str设置成“2014-01-01 00:00:10”
func (r *TaobaoQianniuTasksGetRequest) SetModifyStartTimeStr(_modifyStartTimeStr string) error {
    r._modifyStartTimeStr = _modifyStartTimeStr
    r.Set("modify_start_time_str", _modifyStartTimeStr)
    return nil
}

// ModifyStartTimeStr Getter
func (r TaobaoQianniuTasksGetRequest) GetModifyStartTimeStr() string {
    return r._modifyStartTimeStr
}
// ModifyEndTimeStr Setter
// 根据任务修改时间搜索的结束时间（含），不填则不限。例如查询“2014-01-01 00:00:10”之前有修改的任务，则将modify_end_time_str设置成“2014-01-01 00:00:10”
func (r *TaobaoQianniuTasksGetRequest) SetModifyEndTimeStr(_modifyEndTimeStr string) error {
    r._modifyEndTimeStr = _modifyEndTimeStr
    r.Set("modify_end_time_str", _modifyEndTimeStr)
    return nil
}

// ModifyEndTimeStr Getter
func (r TaobaoQianniuTasksGetRequest) GetModifyEndTimeStr() string {
    return r._modifyEndTimeStr
}
// Priority Setter
// 优先级。即创建时的metadata中的优先级。0为低，1为中，2为高。
func (r *TaobaoQianniuTasksGetRequest) SetPriority(_priority int64) error {
    r._priority = _priority
    r.Set("priority", _priority)
    return nil
}

// Priority Getter
func (r TaobaoQianniuTasksGetRequest) GetPriority() int64 {
    return r._priority
}
// ExcludeBizType Setter
// 需要排除的任务类型
func (r *TaobaoQianniuTasksGetRequest) SetExcludeBizType(_excludeBizType string) error {
    r._excludeBizType = _excludeBizType
    r.Set("exclude_biz_type", _excludeBizType)
    return nil
}

// ExcludeBizType Getter
func (r TaobaoQianniuTasksGetRequest) GetExcludeBizType() string {
    return r._excludeBizType
}
// KeyWord Setter
// 关键词搜索。只对任务内容进行模糊匹配，以及bizid和biznick进行精准匹配
func (r *TaobaoQianniuTasksGetRequest) SetKeyWord(_keyWord string) error {
    r._keyWord = _keyWord
    r.Set("key_word", _keyWord)
    return nil
}

// KeyWord Getter
func (r TaobaoQianniuTasksGetRequest) GetKeyWord() string {
    return r._keyWord
}
// CurrentPage Setter
// 当前页数，从1开始
func (r *TaobaoQianniuTasksGetRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r TaobaoQianniuTasksGetRequest) GetCurrentPage() int64 {
    return r._currentPage
}
// PageSize Setter
// 每页条数
func (r *TaobaoQianniuTasksGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoQianniuTasksGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// BizType Setter
// 业务类型
func (r *TaobaoQianniuTasksGetRequest) SetBizType(_bizType string) error {
    r._bizType = _bizType
    r.Set("biz_type", _bizType)
    return nil
}

// BizType Getter
func (r TaobaoQianniuTasksGetRequest) GetBizType() string {
    return r._bizType
}
// SubBizType Setter
// 子任务类型
func (r *TaobaoQianniuTasksGetRequest) SetSubBizType(_subBizType string) error {
    r._subBizType = _subBizType
    r.Set("sub_biz_type", _subBizType)
    return nil
}

// SubBizType Getter
func (r TaobaoQianniuTasksGetRequest) GetSubBizType() string {
    return r._subBizType
}
// TaskIds Setter
// 任务的ID列表，用逗号分隔
func (r *TaobaoQianniuTasksGetRequest) SetTaskIds(_taskIds string) error {
    r._taskIds = _taskIds
    r.Set("task_ids", _taskIds)
    return nil
}

// TaskIds Getter
func (r TaobaoQianniuTasksGetRequest) GetTaskIds() string {
    return r._taskIds
}
// BizIds Setter
// 业务ID列表，逗号分隔
func (r *TaobaoQianniuTasksGetRequest) SetBizIds(_bizIds string) error {
    r._bizIds = _bizIds
    r.Set("biz_ids", _bizIds)
    return nil
}

// BizIds Getter
func (r TaobaoQianniuTasksGetRequest) GetBizIds() string {
    return r._bizIds
}
// ReceiverUid Setter
// 任务执行者用户数字ID
func (r *TaobaoQianniuTasksGetRequest) SetReceiverUid(_receiverUid int64) error {
    r._receiverUid = _receiverUid
    r.Set("receiver_uid", _receiverUid)
    return nil
}

// ReceiverUid Getter
func (r TaobaoQianniuTasksGetRequest) GetReceiverUid() int64 {
    return r._receiverUid
}
// SenderUid Setter
// 任务发起者用户数字ID
func (r *TaobaoQianniuTasksGetRequest) SetSenderUid(_senderUid int64) error {
    r._senderUid = _senderUid
    r.Set("sender_uid", _senderUid)
    return nil
}

// SenderUid Getter
func (r TaobaoQianniuTasksGetRequest) GetSenderUid() int64 {
    return r._senderUid
}
// Status Setter
// 逗号分隔的任务状态：0-未执行，1-执行中，2-执行完成，3-超时，4-取消，5-忽略
func (r *TaobaoQianniuTasksGetRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoQianniuTasksGetRequest) GetStatus() string {
    return r._status
}
// SubStatus Setter
// 逗号分隔的子任务状态，由业务方自定义
func (r *TaobaoQianniuTasksGetRequest) SetSubStatus(_subStatus string) error {
    r._subStatus = _subStatus
    r.Set("sub_status", _subStatus)
    return nil
}

// SubStatus Getter
func (r TaobaoQianniuTasksGetRequest) GetSubStatus() string {
    return r._subStatus
}
// MetadataIds Setter
// 任务元id，多个以逗号分隔
func (r *TaobaoQianniuTasksGetRequest) SetMetadataIds(_metadataIds string) error {
    r._metadataIds = _metadataIds
    r.Set("metadata_ids", _metadataIds)
    return nil
}

// MetadataIds Getter
func (r TaobaoQianniuTasksGetRequest) GetMetadataIds() string {
    return r._metadataIds
}
// NeedMeta Setter
// 是否需要meta信息，默认值为false
func (r *TaobaoQianniuTasksGetRequest) SetNeedMeta(_needMeta bool) error {
    r._needMeta = _needMeta
    r.Set("need_meta", _needMeta)
    return nil
}

// NeedMeta Getter
func (r TaobaoQianniuTasksGetRequest) GetNeedMeta() bool {
    return r._needMeta
}
// Fields Setter
// 逗号分隔的字段列表，各个字段含义： id：任务ID receiver_uid：执行者用户数字ID receiver_nick：执行者用户昵称 status：任务状态：0-未执行，1-执行中，2-执行完成，3-超时，4-取消，5-忽略 sub_status：子任务状态，由业务方自定义 finish_strategy：任务完成策略：1-一个人完成，2-所有人完成 gmt_finished：任务完成时间，格式：时间毫秒数 biz_type：业务类型 sub_biz_type：子业务类型 biz_id：业务ID biz_param：业务参数 biz_entry：业务入口 tag：任务标签 memo：任务备注
func (r *TaobaoQianniuTasksGetRequest) SetFields(_fields string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TaobaoQianniuTasksGetRequest) GetFields() string {
    return r._fields
}
// ClientInfo Setter
// 客户端的版本信息
func (r *TaobaoQianniuTasksGetRequest) SetClientInfo(_clientInfo string) error {
    r._clientInfo = _clientInfo
    r.Set("client_info", _clientInfo)
    return nil
}

// ClientInfo Getter
func (r TaobaoQianniuTasksGetRequest) GetClientInfo() string {
    return r._clientInfo
}
// NeedDeleted Setter
// 是否需要删除的任务，默认为false
func (r *TaobaoQianniuTasksGetRequest) SetNeedDeleted(_needDeleted bool) error {
    r._needDeleted = _needDeleted
    r.Set("need_deleted", _needDeleted)
    return nil
}

// NeedDeleted Getter
func (r TaobaoQianniuTasksGetRequest) GetNeedDeleted() bool {
    return r._needDeleted
}
