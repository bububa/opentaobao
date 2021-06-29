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
    orderBy   string
    // asc为升，desc为降
    orderType   string
    // 0-不需要提醒，未设提醒时间 1-设置过提醒时间，需要提醒
    remindFlag   int64
    // 业务相关的对象，当前主要表示买家nick
    bizNick   string
    // 根据任务创建时间搜索的开始日期（含），不填则不限。例如只查询2014-01-01当天的任务，则将start_date和end_date都设置成2014-01-01
    startDate   string
    // 根据任务创建时间搜索的结束日期（含），不填则不限。例如只查询2014-01-01当天的任务，则将start_date和end_date都设置成2014-01-01
    endDate   string
    // 根据任务修改时间搜索的开始时间（含），不填则不限。例如查询“2014-01-01 00:00:10”之后有修改的任务，则将modify_start_time_str设置成“2014-01-01 00:00:10”
    modifyStartTimeStr   string
    // 根据任务修改时间搜索的结束时间（含），不填则不限。例如查询“2014-01-01 00:00:10”之前有修改的任务，则将modify_end_time_str设置成“2014-01-01 00:00:10”
    modifyEndTimeStr   string
    // 优先级。即创建时的metadata中的优先级。0为低，1为中，2为高。
    priority   int64
    // 需要排除的任务类型
    excludeBizType   string
    // 关键词搜索。只对任务内容进行模糊匹配，以及bizid和biznick进行精准匹配
    keyWord   string
    // 当前页数，从1开始
    currentPage   int64
    // 每页条数
    pageSize   int64
    // 业务类型
    bizType   string
    // 子任务类型
    subBizType   string
    // 任务的ID列表，用逗号分隔
    taskIds   string
    // 业务ID列表，逗号分隔
    bizIds   string
    // 任务执行者用户数字ID
    receiverUid   int64
    // 任务发起者用户数字ID
    senderUid   int64
    // 逗号分隔的任务状态：0-未执行，1-执行中，2-执行完成，3-超时，4-取消，5-忽略
    status   string
    // 逗号分隔的子任务状态，由业务方自定义
    subStatus   string
    // 任务元id，多个以逗号分隔
    metadataIds   string
    // 是否需要meta信息，默认值为false
    needMeta   bool
    // 逗号分隔的字段列表，各个字段含义： id：任务ID receiver_uid：执行者用户数字ID receiver_nick：执行者用户昵称 status：任务状态：0-未执行，1-执行中，2-执行完成，3-超时，4-取消，5-忽略 sub_status：子任务状态，由业务方自定义 finish_strategy：任务完成策略：1-一个人完成，2-所有人完成 gmt_finished：任务完成时间，格式：时间毫秒数 biz_type：业务类型 sub_biz_type：子业务类型 biz_id：业务ID biz_param：业务参数 biz_entry：业务入口 tag：任务标签 memo：任务备注
    fields   string
    // 客户端的版本信息
    clientInfo   string
    // 是否需要删除的任务，默认为false
    needDeleted   bool
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
func (r *TaobaoQianniuTasksGetRequest) SetOrderBy(orderBy string) error {
    r.orderBy = orderBy
    r.Set("order_by", orderBy)
    return nil
}

// OrderBy Getter
func (r TaobaoQianniuTasksGetRequest) GetOrderBy() string {
    return r.orderBy
}
// OrderType Setter
// asc为升，desc为降
func (r *TaobaoQianniuTasksGetRequest) SetOrderType(orderType string) error {
    r.orderType = orderType
    r.Set("order_type", orderType)
    return nil
}

// OrderType Getter
func (r TaobaoQianniuTasksGetRequest) GetOrderType() string {
    return r.orderType
}
// RemindFlag Setter
// 0-不需要提醒，未设提醒时间 1-设置过提醒时间，需要提醒
func (r *TaobaoQianniuTasksGetRequest) SetRemindFlag(remindFlag int64) error {
    r.remindFlag = remindFlag
    r.Set("remind_flag", remindFlag)
    return nil
}

// RemindFlag Getter
func (r TaobaoQianniuTasksGetRequest) GetRemindFlag() int64 {
    return r.remindFlag
}
// BizNick Setter
// 业务相关的对象，当前主要表示买家nick
func (r *TaobaoQianniuTasksGetRequest) SetBizNick(bizNick string) error {
    r.bizNick = bizNick
    r.Set("biz_nick", bizNick)
    return nil
}

// BizNick Getter
func (r TaobaoQianniuTasksGetRequest) GetBizNick() string {
    return r.bizNick
}
// StartDate Setter
// 根据任务创建时间搜索的开始日期（含），不填则不限。例如只查询2014-01-01当天的任务，则将start_date和end_date都设置成2014-01-01
func (r *TaobaoQianniuTasksGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

// StartDate Getter
func (r TaobaoQianniuTasksGetRequest) GetStartDate() string {
    return r.startDate
}
// EndDate Setter
// 根据任务创建时间搜索的结束日期（含），不填则不限。例如只查询2014-01-01当天的任务，则将start_date和end_date都设置成2014-01-01
func (r *TaobaoQianniuTasksGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r TaobaoQianniuTasksGetRequest) GetEndDate() string {
    return r.endDate
}
// ModifyStartTimeStr Setter
// 根据任务修改时间搜索的开始时间（含），不填则不限。例如查询“2014-01-01 00:00:10”之后有修改的任务，则将modify_start_time_str设置成“2014-01-01 00:00:10”
func (r *TaobaoQianniuTasksGetRequest) SetModifyStartTimeStr(modifyStartTimeStr string) error {
    r.modifyStartTimeStr = modifyStartTimeStr
    r.Set("modify_start_time_str", modifyStartTimeStr)
    return nil
}

// ModifyStartTimeStr Getter
func (r TaobaoQianniuTasksGetRequest) GetModifyStartTimeStr() string {
    return r.modifyStartTimeStr
}
// ModifyEndTimeStr Setter
// 根据任务修改时间搜索的结束时间（含），不填则不限。例如查询“2014-01-01 00:00:10”之前有修改的任务，则将modify_end_time_str设置成“2014-01-01 00:00:10”
func (r *TaobaoQianniuTasksGetRequest) SetModifyEndTimeStr(modifyEndTimeStr string) error {
    r.modifyEndTimeStr = modifyEndTimeStr
    r.Set("modify_end_time_str", modifyEndTimeStr)
    return nil
}

// ModifyEndTimeStr Getter
func (r TaobaoQianniuTasksGetRequest) GetModifyEndTimeStr() string {
    return r.modifyEndTimeStr
}
// Priority Setter
// 优先级。即创建时的metadata中的优先级。0为低，1为中，2为高。
func (r *TaobaoQianniuTasksGetRequest) SetPriority(priority int64) error {
    r.priority = priority
    r.Set("priority", priority)
    return nil
}

// Priority Getter
func (r TaobaoQianniuTasksGetRequest) GetPriority() int64 {
    return r.priority
}
// ExcludeBizType Setter
// 需要排除的任务类型
func (r *TaobaoQianniuTasksGetRequest) SetExcludeBizType(excludeBizType string) error {
    r.excludeBizType = excludeBizType
    r.Set("exclude_biz_type", excludeBizType)
    return nil
}

// ExcludeBizType Getter
func (r TaobaoQianniuTasksGetRequest) GetExcludeBizType() string {
    return r.excludeBizType
}
// KeyWord Setter
// 关键词搜索。只对任务内容进行模糊匹配，以及bizid和biznick进行精准匹配
func (r *TaobaoQianniuTasksGetRequest) SetKeyWord(keyWord string) error {
    r.keyWord = keyWord
    r.Set("key_word", keyWord)
    return nil
}

// KeyWord Getter
func (r TaobaoQianniuTasksGetRequest) GetKeyWord() string {
    return r.keyWord
}
// CurrentPage Setter
// 当前页数，从1开始
func (r *TaobaoQianniuTasksGetRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

// CurrentPage Getter
func (r TaobaoQianniuTasksGetRequest) GetCurrentPage() int64 {
    return r.currentPage
}
// PageSize Setter
// 每页条数
func (r *TaobaoQianniuTasksGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoQianniuTasksGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// BizType Setter
// 业务类型
func (r *TaobaoQianniuTasksGetRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

// BizType Getter
func (r TaobaoQianniuTasksGetRequest) GetBizType() string {
    return r.bizType
}
// SubBizType Setter
// 子任务类型
func (r *TaobaoQianniuTasksGetRequest) SetSubBizType(subBizType string) error {
    r.subBizType = subBizType
    r.Set("sub_biz_type", subBizType)
    return nil
}

// SubBizType Getter
func (r TaobaoQianniuTasksGetRequest) GetSubBizType() string {
    return r.subBizType
}
// TaskIds Setter
// 任务的ID列表，用逗号分隔
func (r *TaobaoQianniuTasksGetRequest) SetTaskIds(taskIds string) error {
    r.taskIds = taskIds
    r.Set("task_ids", taskIds)
    return nil
}

// TaskIds Getter
func (r TaobaoQianniuTasksGetRequest) GetTaskIds() string {
    return r.taskIds
}
// BizIds Setter
// 业务ID列表，逗号分隔
func (r *TaobaoQianniuTasksGetRequest) SetBizIds(bizIds string) error {
    r.bizIds = bizIds
    r.Set("biz_ids", bizIds)
    return nil
}

// BizIds Getter
func (r TaobaoQianniuTasksGetRequest) GetBizIds() string {
    return r.bizIds
}
// ReceiverUid Setter
// 任务执行者用户数字ID
func (r *TaobaoQianniuTasksGetRequest) SetReceiverUid(receiverUid int64) error {
    r.receiverUid = receiverUid
    r.Set("receiver_uid", receiverUid)
    return nil
}

// ReceiverUid Getter
func (r TaobaoQianniuTasksGetRequest) GetReceiverUid() int64 {
    return r.receiverUid
}
// SenderUid Setter
// 任务发起者用户数字ID
func (r *TaobaoQianniuTasksGetRequest) SetSenderUid(senderUid int64) error {
    r.senderUid = senderUid
    r.Set("sender_uid", senderUid)
    return nil
}

// SenderUid Getter
func (r TaobaoQianniuTasksGetRequest) GetSenderUid() int64 {
    return r.senderUid
}
// Status Setter
// 逗号分隔的任务状态：0-未执行，1-执行中，2-执行完成，3-超时，4-取消，5-忽略
func (r *TaobaoQianniuTasksGetRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoQianniuTasksGetRequest) GetStatus() string {
    return r.status
}
// SubStatus Setter
// 逗号分隔的子任务状态，由业务方自定义
func (r *TaobaoQianniuTasksGetRequest) SetSubStatus(subStatus string) error {
    r.subStatus = subStatus
    r.Set("sub_status", subStatus)
    return nil
}

// SubStatus Getter
func (r TaobaoQianniuTasksGetRequest) GetSubStatus() string {
    return r.subStatus
}
// MetadataIds Setter
// 任务元id，多个以逗号分隔
func (r *TaobaoQianniuTasksGetRequest) SetMetadataIds(metadataIds string) error {
    r.metadataIds = metadataIds
    r.Set("metadata_ids", metadataIds)
    return nil
}

// MetadataIds Getter
func (r TaobaoQianniuTasksGetRequest) GetMetadataIds() string {
    return r.metadataIds
}
// NeedMeta Setter
// 是否需要meta信息，默认值为false
func (r *TaobaoQianniuTasksGetRequest) SetNeedMeta(needMeta bool) error {
    r.needMeta = needMeta
    r.Set("need_meta", needMeta)
    return nil
}

// NeedMeta Getter
func (r TaobaoQianniuTasksGetRequest) GetNeedMeta() bool {
    return r.needMeta
}
// Fields Setter
// 逗号分隔的字段列表，各个字段含义： id：任务ID receiver_uid：执行者用户数字ID receiver_nick：执行者用户昵称 status：任务状态：0-未执行，1-执行中，2-执行完成，3-超时，4-取消，5-忽略 sub_status：子任务状态，由业务方自定义 finish_strategy：任务完成策略：1-一个人完成，2-所有人完成 gmt_finished：任务完成时间，格式：时间毫秒数 biz_type：业务类型 sub_biz_type：子业务类型 biz_id：业务ID biz_param：业务参数 biz_entry：业务入口 tag：任务标签 memo：任务备注
func (r *TaobaoQianniuTasksGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r TaobaoQianniuTasksGetRequest) GetFields() string {
    return r.fields
}
// ClientInfo Setter
// 客户端的版本信息
func (r *TaobaoQianniuTasksGetRequest) SetClientInfo(clientInfo string) error {
    r.clientInfo = clientInfo
    r.Set("client_info", clientInfo)
    return nil
}

// ClientInfo Getter
func (r TaobaoQianniuTasksGetRequest) GetClientInfo() string {
    return r.clientInfo
}
// NeedDeleted Setter
// 是否需要删除的任务，默认为false
func (r *TaobaoQianniuTasksGetRequest) SetNeedDeleted(needDeleted bool) error {
    r.needDeleted = needDeleted
    r.Set("need_deleted", needDeleted)
    return nil
}

// NeedDeleted Getter
func (r TaobaoQianniuTasksGetRequest) GetNeedDeleted() bool {
    return r.needDeleted
}
