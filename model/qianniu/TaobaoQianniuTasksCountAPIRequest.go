package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQianniuTasksCountAPIRequest 任务查询条数接口 API请求
// taobao.qianniu.tasks.count
//
// 任务查询条数接口
type TaobaoQianniuTasksCountAPIRequest struct {
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
	// 与业务相关的买家nick
	_bizNick string
	// 按时间段搜索时的开始日期，格式如2014-01-01，不填则不限
	_startDate string
	// 按时间段搜索的结束日期。不填则不限。格式为2014-01-01
	_endDate string
	// 需要排除的任务类型
	_excludeBizType string
	// 关键词搜索。只对任务内容进行模糊匹配，以及bizid和biznick进行精准匹配
	_keyWord string
	// 任务执行者用户数字ID
	_receiverUid int64
	// 任务发起者用户数字ID
	_senderUid int64
	// 0-不需要提醒，未设提醒时间 1-设置过提醒时间，需要提醒
	_remindFlag int64
	// 优先级
	_priority int64
}

// NewTaobaoQianniuTasksCountRequest 初始化TaobaoQianniuTasksCountAPIRequest对象
func NewTaobaoQianniuTasksCountRequest() *TaobaoQianniuTasksCountAPIRequest {
	return &TaobaoQianniuTasksCountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQianniuTasksCountAPIRequest) GetApiMethodName() string {
	return "taobao.qianniu.tasks.count"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQianniuTasksCountAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBizType is BizType Setter
// 业务类型
func (r *TaobaoQianniuTasksCountAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TaobaoQianniuTasksCountAPIRequest) GetBizType() string {
	return r._bizType
}

// SetSubBizType is SubBizType Setter
// 子任务类型
func (r *TaobaoQianniuTasksCountAPIRequest) SetSubBizType(_subBizType string) error {
	r._subBizType = _subBizType
	r.Set("sub_biz_type", _subBizType)
	return nil
}

// GetSubBizType SubBizType Getter
func (r TaobaoQianniuTasksCountAPIRequest) GetSubBizType() string {
	return r._subBizType
}

// SetTaskIds is TaskIds Setter
// 任务的ID列表，用逗号分隔
func (r *TaobaoQianniuTasksCountAPIRequest) SetTaskIds(_taskIds string) error {
	r._taskIds = _taskIds
	r.Set("task_ids", _taskIds)
	return nil
}

// GetTaskIds TaskIds Getter
func (r TaobaoQianniuTasksCountAPIRequest) GetTaskIds() string {
	return r._taskIds
}

// SetBizIds is BizIds Setter
// 业务ID列表，逗号分隔
func (r *TaobaoQianniuTasksCountAPIRequest) SetBizIds(_bizIds string) error {
	r._bizIds = _bizIds
	r.Set("biz_ids", _bizIds)
	return nil
}

// GetBizIds BizIds Getter
func (r TaobaoQianniuTasksCountAPIRequest) GetBizIds() string {
	return r._bizIds
}

// SetStatus is Status Setter
// 逗号分隔的任务状态：0-未执行，1-执行中，2-执行完成，3-超时，4-取消，5-忽略
func (r *TaobaoQianniuTasksCountAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoQianniuTasksCountAPIRequest) GetStatus() string {
	return r._status
}

// SetSubStatus is SubStatus Setter
// 逗号分隔的子任务状态，由业务方自定义
func (r *TaobaoQianniuTasksCountAPIRequest) SetSubStatus(_subStatus string) error {
	r._subStatus = _subStatus
	r.Set("sub_status", _subStatus)
	return nil
}

// GetSubStatus SubStatus Getter
func (r TaobaoQianniuTasksCountAPIRequest) GetSubStatus() string {
	return r._subStatus
}

// SetMetadataIds is MetadataIds Setter
// 任务元id，多个以逗号分隔
func (r *TaobaoQianniuTasksCountAPIRequest) SetMetadataIds(_metadataIds string) error {
	r._metadataIds = _metadataIds
	r.Set("metadata_ids", _metadataIds)
	return nil
}

// GetMetadataIds MetadataIds Getter
func (r TaobaoQianniuTasksCountAPIRequest) GetMetadataIds() string {
	return r._metadataIds
}

// SetBizNick is BizNick Setter
// 与业务相关的买家nick
func (r *TaobaoQianniuTasksCountAPIRequest) SetBizNick(_bizNick string) error {
	r._bizNick = _bizNick
	r.Set("biz_nick", _bizNick)
	return nil
}

// GetBizNick BizNick Getter
func (r TaobaoQianniuTasksCountAPIRequest) GetBizNick() string {
	return r._bizNick
}

// SetStartDate is StartDate Setter
// 按时间段搜索时的开始日期，格式如2014-01-01，不填则不限
func (r *TaobaoQianniuTasksCountAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaoQianniuTasksCountAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 按时间段搜索的结束日期。不填则不限。格式为2014-01-01
func (r *TaobaoQianniuTasksCountAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaoQianniuTasksCountAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetExcludeBizType is ExcludeBizType Setter
// 需要排除的任务类型
func (r *TaobaoQianniuTasksCountAPIRequest) SetExcludeBizType(_excludeBizType string) error {
	r._excludeBizType = _excludeBizType
	r.Set("exclude_biz_type", _excludeBizType)
	return nil
}

// GetExcludeBizType ExcludeBizType Getter
func (r TaobaoQianniuTasksCountAPIRequest) GetExcludeBizType() string {
	return r._excludeBizType
}

// SetKeyWord is KeyWord Setter
// 关键词搜索。只对任务内容进行模糊匹配，以及bizid和biznick进行精准匹配
func (r *TaobaoQianniuTasksCountAPIRequest) SetKeyWord(_keyWord string) error {
	r._keyWord = _keyWord
	r.Set("key_word", _keyWord)
	return nil
}

// GetKeyWord KeyWord Getter
func (r TaobaoQianniuTasksCountAPIRequest) GetKeyWord() string {
	return r._keyWord
}

// SetReceiverUid is ReceiverUid Setter
// 任务执行者用户数字ID
func (r *TaobaoQianniuTasksCountAPIRequest) SetReceiverUid(_receiverUid int64) error {
	r._receiverUid = _receiverUid
	r.Set("receiver_uid", _receiverUid)
	return nil
}

// GetReceiverUid ReceiverUid Getter
func (r TaobaoQianniuTasksCountAPIRequest) GetReceiverUid() int64 {
	return r._receiverUid
}

// SetSenderUid is SenderUid Setter
// 任务发起者用户数字ID
func (r *TaobaoQianniuTasksCountAPIRequest) SetSenderUid(_senderUid int64) error {
	r._senderUid = _senderUid
	r.Set("sender_uid", _senderUid)
	return nil
}

// GetSenderUid SenderUid Getter
func (r TaobaoQianniuTasksCountAPIRequest) GetSenderUid() int64 {
	return r._senderUid
}

// SetRemindFlag is RemindFlag Setter
// 0-不需要提醒，未设提醒时间 1-设置过提醒时间，需要提醒
func (r *TaobaoQianniuTasksCountAPIRequest) SetRemindFlag(_remindFlag int64) error {
	r._remindFlag = _remindFlag
	r.Set("remind_flag", _remindFlag)
	return nil
}

// GetRemindFlag RemindFlag Getter
func (r TaobaoQianniuTasksCountAPIRequest) GetRemindFlag() int64 {
	return r._remindFlag
}

// SetPriority is Priority Setter
// 优先级
func (r *TaobaoQianniuTasksCountAPIRequest) SetPriority(_priority int64) error {
	r._priority = _priority
	r.Set("priority", _priority)
	return nil
}

// GetPriority Priority Getter
func (r TaobaoQianniuTasksCountAPIRequest) GetPriority() int64 {
	return r._priority
}
