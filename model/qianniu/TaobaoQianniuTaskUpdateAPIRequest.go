package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqianniutaskupdateAPIRequest 更新轻任务 API请求
// taobao.qianniu.task.update
//
// 由任务执行者调用，sub_status，tag和memo至少提供一个
type TaobaoqianniutaskupdateAPIRequest struct {
	model.Params
	// 子任务状态，由业务方自定义
	_subStatus string
	// 状态值，多个以逗号分隔
	_status string
	// 任务标签
	_tag string
	// 任务备注。当memo_mode为1时，memo将采用追加方式。
	_memo string
	// 应用自定义参数
	_bizParam string
	// 任务ID
	_taskId int64
	// 提醒时间，时间的毫秒数
	_remindTime int64
	// 0为不提醒，1为全部提醒，2为PC提醒，3为移动提醒，4为已提醒，5为已忽略。
	_remindFlag int64
	// 表示memo字段的更新策略。如需采用追加方式的，请将此字段设置为1。
	_memoMode int64
	// 默认填0，数字越大优化级越高。当前常用0和1.
	_priority int64
	// 0表示没有删除，1表示删除
	_isDeleted int64
}

// NewTaobaoqianniutaskupdateRequest 初始化TaobaoqianniutaskupdateAPIRequest对象
func NewTaobaoqianniutaskupdateRequest() *TaobaoqianniutaskupdateAPIRequest {
	return &TaobaoqianniutaskupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqianniutaskupdateAPIRequest) GetApiMethodName() string {
	return "taobao.qianniu.task.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqianniutaskupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqianniutaskupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubStatus is SubStatus Setter
// 子任务状态，由业务方自定义
func (r *TaobaoqianniutaskupdateAPIRequest) SetSubStatus(_subStatus string) error {
	r._subStatus = _subStatus
	r.Set("sub_status", _subStatus)
	return nil
}

// GetSubStatus SubStatus Getter
func (r TaobaoqianniutaskupdateAPIRequest) GetSubStatus() string {
	return r._subStatus
}

// SetStatus is Status Setter
// 状态值，多个以逗号分隔
func (r *TaobaoqianniutaskupdateAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoqianniutaskupdateAPIRequest) GetStatus() string {
	return r._status
}

// SetTag is Tag Setter
// 任务标签
func (r *TaobaoqianniutaskupdateAPIRequest) SetTag(_tag string) error {
	r._tag = _tag
	r.Set("tag", _tag)
	return nil
}

// GetTag Tag Getter
func (r TaobaoqianniutaskupdateAPIRequest) GetTag() string {
	return r._tag
}

// SetMemo is Memo Setter
// 任务备注。当memo_mode为1时，memo将采用追加方式。
func (r *TaobaoqianniutaskupdateAPIRequest) SetMemo(_memo string) error {
	r._memo = _memo
	r.Set("memo", _memo)
	return nil
}

// GetMemo Memo Getter
func (r TaobaoqianniutaskupdateAPIRequest) GetMemo() string {
	return r._memo
}

// SetBizParam is BizParam Setter
// 应用自定义参数
func (r *TaobaoqianniutaskupdateAPIRequest) SetBizParam(_bizParam string) error {
	r._bizParam = _bizParam
	r.Set("biz_param", _bizParam)
	return nil
}

// GetBizParam BizParam Getter
func (r TaobaoqianniutaskupdateAPIRequest) GetBizParam() string {
	return r._bizParam
}

// SetTaskId is TaskId Setter
// 任务ID
func (r *TaobaoqianniutaskupdateAPIRequest) SetTaskId(_taskId int64) error {
	r._taskId = _taskId
	r.Set("task_id", _taskId)
	return nil
}

// GetTaskId TaskId Getter
func (r TaobaoqianniutaskupdateAPIRequest) GetTaskId() int64 {
	return r._taskId
}

// SetRemindTime is RemindTime Setter
// 提醒时间，时间的毫秒数
func (r *TaobaoqianniutaskupdateAPIRequest) SetRemindTime(_remindTime int64) error {
	r._remindTime = _remindTime
	r.Set("remind_time", _remindTime)
	return nil
}

// GetRemindTime RemindTime Getter
func (r TaobaoqianniutaskupdateAPIRequest) GetRemindTime() int64 {
	return r._remindTime
}

// SetRemindFlag is RemindFlag Setter
// 0为不提醒，1为全部提醒，2为PC提醒，3为移动提醒，4为已提醒，5为已忽略。
func (r *TaobaoqianniutaskupdateAPIRequest) SetRemindFlag(_remindFlag int64) error {
	r._remindFlag = _remindFlag
	r.Set("remind_flag", _remindFlag)
	return nil
}

// GetRemindFlag RemindFlag Getter
func (r TaobaoqianniutaskupdateAPIRequest) GetRemindFlag() int64 {
	return r._remindFlag
}

// SetMemoMode is MemoMode Setter
// 表示memo字段的更新策略。如需采用追加方式的，请将此字段设置为1。
func (r *TaobaoqianniutaskupdateAPIRequest) SetMemoMode(_memoMode int64) error {
	r._memoMode = _memoMode
	r.Set("memo_mode", _memoMode)
	return nil
}

// GetMemoMode MemoMode Getter
func (r TaobaoqianniutaskupdateAPIRequest) GetMemoMode() int64 {
	return r._memoMode
}

// SetPriority is Priority Setter
// 默认填0，数字越大优化级越高。当前常用0和1.
func (r *TaobaoqianniutaskupdateAPIRequest) SetPriority(_priority int64) error {
	r._priority = _priority
	r.Set("priority", _priority)
	return nil
}

// GetPriority Priority Getter
func (r TaobaoqianniutaskupdateAPIRequest) GetPriority() int64 {
	return r._priority
}

// SetIsDeleted is IsDeleted Setter
// 0表示没有删除，1表示删除
func (r *TaobaoqianniutaskupdateAPIRequest) SetIsDeleted(_isDeleted int64) error {
	r._isDeleted = _isDeleted
	r.Set("is_deleted", _isDeleted)
	return nil
}

// GetIsDeleted IsDeleted Getter
func (r TaobaoqianniutaskupdateAPIRequest) GetIsDeleted() int64 {
	return r._isDeleted
}
