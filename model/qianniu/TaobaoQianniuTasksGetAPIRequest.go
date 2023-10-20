package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqianniutasksgetAPIRequest 获取指定的任务 API请求
// taobao.qianniu.tasks.get
//
// 获取指定的任务，可用的参数组合：&lt;br/&gt;task_ids + need_meta + fields：精确查找&lt;br/&gt;biz_type + sub_biz_type + biz_ids + need_meta + fields：按照业务ID查找&lt;br/&gt;biz_type + sub_biz_type + sender_uid + need_meta + fields：按照发起者查找&lt;br/&gt;biz_type + sub_biz_type + receiver_uid + need_meta + fields：按照执行者查找&lt;br/&gt;biz_type+modify_start_time+modify_end_time+fields:能支持指定修改时间的查询，用于增量查询等
type TaobaoqianniutasksgetAPIRequest struct {
	model.Params
	// 业务类型
	_bizType string
	// 子任务类型
	_subBizType string
	// 任务的ID列表，用逗号分隔
	_taskIds string
	// 业务ID列表，逗号分隔
	_bizIds string
	// 逗号分隔的任务状态：0-未执行，1-执行中，2-执行完成，3-超时，4-取消，5-忽略
	_status string
	// 逗号分隔的子任务状态，由业务方自定义
	_subStatus string
	// 任务元id，多个以逗号分隔
	_metadataIds string
	// 逗号分隔的字段列表，各个字段含义： id：任务ID receiver_uid：执行者用户数字ID receiver_nick：执行者用户昵称 status：任务状态：0-未执行，1-执行中，2-执行完成，3-超时，4-取消，5-忽略 sub_status：子任务状态，由业务方自定义 finish_strategy：任务完成策略：1-一个人完成，2-所有人完成 gmt_finished：任务完成时间，格式：时间毫秒数 biz_type：业务类型 sub_biz_type：子业务类型 biz_id：业务ID biz_param：业务参数 biz_entry：业务入口 tag：任务标签 memo：任务备注
	_fields string
	// 排序字段，可以为id,gmt_create,gmt_finished,metadata_id等
	_orderBy string
	// asc为升，desc为降
	_orderType string
	// 业务相关的对象，当前主要表示买家nick
	_bizNick string
	// 根据任务创建时间搜索的开始日期（含），不填则不限。例如只查询2014-01-01当天的任务，则将start_date和end_date都设置成2014-01-01
	_startDate string
	// 根据任务创建时间搜索的结束日期（含），不填则不限。例如只查询2014-01-01当天的任务，则将start_date和end_date都设置成2014-01-01
	_endDate string
	// 根据任务修改时间搜索的开始时间（含），不填则不限。例如查询“2014-01-01 00:00:10”之后有修改的任务，则将modify_start_time_str设置成“2014-01-01 00:00:10”
	_modifyStartTimeStr string
	// 根据任务修改时间搜索的结束时间（含），不填则不限。例如查询“2014-01-01 00:00:10”之前有修改的任务，则将modify_end_time_str设置成“2014-01-01 00:00:10”
	_modifyEndTimeStr string
	// 需要排除的任务类型
	_excludeBizType string
	// 关键词搜索。只对任务内容进行模糊匹配，以及bizid和biznick进行精准匹配
	_keyWord string
	// 客户端的版本信息
	_clientInfo string
	// 任务执行者用户数字ID
	_receiverUid int64
	// 任务发起者用户数字ID
	_senderUid int64
	// 每页条数
	_pageSize int64
	// 当前页数，从1开始
	_currentPage int64
	// 0-不需要提醒，未设提醒时间 1-设置过提醒时间，需要提醒
	_remindFlag int64
	// 优先级。即创建时的metadata中的优先级。0为低，1为中，2为高。
	_priority int64
	// 是否需要meta信息，默认值为false
	_needMeta bool
	// 是否需要删除的任务，默认为false
	_needDeleted bool
}

// NewTaobaoqianniutasksgetRequest 初始化TaobaoqianniutasksgetAPIRequest对象
func NewTaobaoqianniutasksgetRequest() *TaobaoqianniutasksgetAPIRequest {
	return &TaobaoqianniutasksgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqianniutasksgetAPIRequest) GetApiMethodName() string {
	return "taobao.qianniu.tasks.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqianniutasksgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqianniutasksgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizType is BizType Setter
// 业务类型
func (r *TaobaoqianniutasksgetAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TaobaoqianniutasksgetAPIRequest) GetBizType() string {
	return r._bizType
}

// SetSubBizType is SubBizType Setter
// 子任务类型
func (r *TaobaoqianniutasksgetAPIRequest) SetSubBizType(_subBizType string) error {
	r._subBizType = _subBizType
	r.Set("sub_biz_type", _subBizType)
	return nil
}

// GetSubBizType SubBizType Getter
func (r TaobaoqianniutasksgetAPIRequest) GetSubBizType() string {
	return r._subBizType
}

// SetTaskIds is TaskIds Setter
// 任务的ID列表，用逗号分隔
func (r *TaobaoqianniutasksgetAPIRequest) SetTaskIds(_taskIds string) error {
	r._taskIds = _taskIds
	r.Set("task_ids", _taskIds)
	return nil
}

// GetTaskIds TaskIds Getter
func (r TaobaoqianniutasksgetAPIRequest) GetTaskIds() string {
	return r._taskIds
}

// SetBizIds is BizIds Setter
// 业务ID列表，逗号分隔
func (r *TaobaoqianniutasksgetAPIRequest) SetBizIds(_bizIds string) error {
	r._bizIds = _bizIds
	r.Set("biz_ids", _bizIds)
	return nil
}

// GetBizIds BizIds Getter
func (r TaobaoqianniutasksgetAPIRequest) GetBizIds() string {
	return r._bizIds
}

// SetStatus is Status Setter
// 逗号分隔的任务状态：0-未执行，1-执行中，2-执行完成，3-超时，4-取消，5-忽略
func (r *TaobaoqianniutasksgetAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoqianniutasksgetAPIRequest) GetStatus() string {
	return r._status
}

// SetSubStatus is SubStatus Setter
// 逗号分隔的子任务状态，由业务方自定义
func (r *TaobaoqianniutasksgetAPIRequest) SetSubStatus(_subStatus string) error {
	r._subStatus = _subStatus
	r.Set("sub_status", _subStatus)
	return nil
}

// GetSubStatus SubStatus Getter
func (r TaobaoqianniutasksgetAPIRequest) GetSubStatus() string {
	return r._subStatus
}

// SetMetadataIds is MetadataIds Setter
// 任务元id，多个以逗号分隔
func (r *TaobaoqianniutasksgetAPIRequest) SetMetadataIds(_metadataIds string) error {
	r._metadataIds = _metadataIds
	r.Set("metadata_ids", _metadataIds)
	return nil
}

// GetMetadataIds MetadataIds Getter
func (r TaobaoqianniutasksgetAPIRequest) GetMetadataIds() string {
	return r._metadataIds
}

// SetFields is Fields Setter
// 逗号分隔的字段列表，各个字段含义： id：任务ID receiver_uid：执行者用户数字ID receiver_nick：执行者用户昵称 status：任务状态：0-未执行，1-执行中，2-执行完成，3-超时，4-取消，5-忽略 sub_status：子任务状态，由业务方自定义 finish_strategy：任务完成策略：1-一个人完成，2-所有人完成 gmt_finished：任务完成时间，格式：时间毫秒数 biz_type：业务类型 sub_biz_type：子业务类型 biz_id：业务ID biz_param：业务参数 biz_entry：业务入口 tag：任务标签 memo：任务备注
func (r *TaobaoqianniutasksgetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoqianniutasksgetAPIRequest) GetFields() string {
	return r._fields
}

// SetOrderBy is OrderBy Setter
// 排序字段，可以为id,gmt_create,gmt_finished,metadata_id等
func (r *TaobaoqianniutasksgetAPIRequest) SetOrderBy(_orderBy string) error {
	r._orderBy = _orderBy
	r.Set("order_by", _orderBy)
	return nil
}

// GetOrderBy OrderBy Getter
func (r TaobaoqianniutasksgetAPIRequest) GetOrderBy() string {
	return r._orderBy
}

// SetOrderType is OrderType Setter
// asc为升，desc为降
func (r *TaobaoqianniutasksgetAPIRequest) SetOrderType(_orderType string) error {
	r._orderType = _orderType
	r.Set("order_type", _orderType)
	return nil
}

// GetOrderType OrderType Getter
func (r TaobaoqianniutasksgetAPIRequest) GetOrderType() string {
	return r._orderType
}

// SetBizNick is BizNick Setter
// 业务相关的对象，当前主要表示买家nick
func (r *TaobaoqianniutasksgetAPIRequest) SetBizNick(_bizNick string) error {
	r._bizNick = _bizNick
	r.Set("biz_nick", _bizNick)
	return nil
}

// GetBizNick BizNick Getter
func (r TaobaoqianniutasksgetAPIRequest) GetBizNick() string {
	return r._bizNick
}

// SetStartDate is StartDate Setter
// 根据任务创建时间搜索的开始日期（含），不填则不限。例如只查询2014-01-01当天的任务，则将start_date和end_date都设置成2014-01-01
func (r *TaobaoqianniutasksgetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaoqianniutasksgetAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 根据任务创建时间搜索的结束日期（含），不填则不限。例如只查询2014-01-01当天的任务，则将start_date和end_date都设置成2014-01-01
func (r *TaobaoqianniutasksgetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaoqianniutasksgetAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetModifyStartTimeStr is ModifyStartTimeStr Setter
// 根据任务修改时间搜索的开始时间（含），不填则不限。例如查询“2014-01-01 00:00:10”之后有修改的任务，则将modify_start_time_str设置成“2014-01-01 00:00:10”
func (r *TaobaoqianniutasksgetAPIRequest) SetModifyStartTimeStr(_modifyStartTimeStr string) error {
	r._modifyStartTimeStr = _modifyStartTimeStr
	r.Set("modify_start_time_str", _modifyStartTimeStr)
	return nil
}

// GetModifyStartTimeStr ModifyStartTimeStr Getter
func (r TaobaoqianniutasksgetAPIRequest) GetModifyStartTimeStr() string {
	return r._modifyStartTimeStr
}

// SetModifyEndTimeStr is ModifyEndTimeStr Setter
// 根据任务修改时间搜索的结束时间（含），不填则不限。例如查询“2014-01-01 00:00:10”之前有修改的任务，则将modify_end_time_str设置成“2014-01-01 00:00:10”
func (r *TaobaoqianniutasksgetAPIRequest) SetModifyEndTimeStr(_modifyEndTimeStr string) error {
	r._modifyEndTimeStr = _modifyEndTimeStr
	r.Set("modify_end_time_str", _modifyEndTimeStr)
	return nil
}

// GetModifyEndTimeStr ModifyEndTimeStr Getter
func (r TaobaoqianniutasksgetAPIRequest) GetModifyEndTimeStr() string {
	return r._modifyEndTimeStr
}

// SetExcludeBizType is ExcludeBizType Setter
// 需要排除的任务类型
func (r *TaobaoqianniutasksgetAPIRequest) SetExcludeBizType(_excludeBizType string) error {
	r._excludeBizType = _excludeBizType
	r.Set("exclude_biz_type", _excludeBizType)
	return nil
}

// GetExcludeBizType ExcludeBizType Getter
func (r TaobaoqianniutasksgetAPIRequest) GetExcludeBizType() string {
	return r._excludeBizType
}

// SetKeyWord is KeyWord Setter
// 关键词搜索。只对任务内容进行模糊匹配，以及bizid和biznick进行精准匹配
func (r *TaobaoqianniutasksgetAPIRequest) SetKeyWord(_keyWord string) error {
	r._keyWord = _keyWord
	r.Set("key_word", _keyWord)
	return nil
}

// GetKeyWord KeyWord Getter
func (r TaobaoqianniutasksgetAPIRequest) GetKeyWord() string {
	return r._keyWord
}

// SetClientInfo is ClientInfo Setter
// 客户端的版本信息
func (r *TaobaoqianniutasksgetAPIRequest) SetClientInfo(_clientInfo string) error {
	r._clientInfo = _clientInfo
	r.Set("client_info", _clientInfo)
	return nil
}

// GetClientInfo ClientInfo Getter
func (r TaobaoqianniutasksgetAPIRequest) GetClientInfo() string {
	return r._clientInfo
}

// SetReceiverUid is ReceiverUid Setter
// 任务执行者用户数字ID
func (r *TaobaoqianniutasksgetAPIRequest) SetReceiverUid(_receiverUid int64) error {
	r._receiverUid = _receiverUid
	r.Set("receiver_uid", _receiverUid)
	return nil
}

// GetReceiverUid ReceiverUid Getter
func (r TaobaoqianniutasksgetAPIRequest) GetReceiverUid() int64 {
	return r._receiverUid
}

// SetSenderUid is SenderUid Setter
// 任务发起者用户数字ID
func (r *TaobaoqianniutasksgetAPIRequest) SetSenderUid(_senderUid int64) error {
	r._senderUid = _senderUid
	r.Set("sender_uid", _senderUid)
	return nil
}

// GetSenderUid SenderUid Getter
func (r TaobaoqianniutasksgetAPIRequest) GetSenderUid() int64 {
	return r._senderUid
}

// SetPageSize is PageSize Setter
// 每页条数
func (r *TaobaoqianniutasksgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoqianniutasksgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetCurrentPage is CurrentPage Setter
// 当前页数，从1开始
func (r *TaobaoqianniutasksgetAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r TaobaoqianniutasksgetAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetRemindFlag is RemindFlag Setter
// 0-不需要提醒，未设提醒时间 1-设置过提醒时间，需要提醒
func (r *TaobaoqianniutasksgetAPIRequest) SetRemindFlag(_remindFlag int64) error {
	r._remindFlag = _remindFlag
	r.Set("remind_flag", _remindFlag)
	return nil
}

// GetRemindFlag RemindFlag Getter
func (r TaobaoqianniutasksgetAPIRequest) GetRemindFlag() int64 {
	return r._remindFlag
}

// SetPriority is Priority Setter
// 优先级。即创建时的metadata中的优先级。0为低，1为中，2为高。
func (r *TaobaoqianniutasksgetAPIRequest) SetPriority(_priority int64) error {
	r._priority = _priority
	r.Set("priority", _priority)
	return nil
}

// GetPriority Priority Getter
func (r TaobaoqianniutasksgetAPIRequest) GetPriority() int64 {
	return r._priority
}

// SetNeedMeta is NeedMeta Setter
// 是否需要meta信息，默认值为false
func (r *TaobaoqianniutasksgetAPIRequest) SetNeedMeta(_needMeta bool) error {
	r._needMeta = _needMeta
	r.Set("need_meta", _needMeta)
	return nil
}

// GetNeedMeta NeedMeta Getter
func (r TaobaoqianniutasksgetAPIRequest) GetNeedMeta() bool {
	return r._needMeta
}

// SetNeedDeleted is NeedDeleted Setter
// 是否需要删除的任务，默认为false
func (r *TaobaoqianniutasksgetAPIRequest) SetNeedDeleted(_needDeleted bool) error {
	r._needDeleted = _needDeleted
	r.Set("need_deleted", _needDeleted)
	return nil
}

// GetNeedDeleted NeedDeleted Getter
func (r TaobaoqianniutasksgetAPIRequest) GetNeedDeleted() bool {
	return r._needDeleted
}
