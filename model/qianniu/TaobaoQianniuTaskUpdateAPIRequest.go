package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQianniuTaskUpdateAPIRequest 更新轻任务 API请求
// taobao.qianniu.task.update
//
// 由任务执行者调用，sub_status，tag和memo至少提供一个
type TaobaoQianniuTaskUpdateAPIRequest struct {
	model.Params
	// 子任务状态，由业务方自定义
	_subStatus string
	// 任务标签
	_tag string
	// 任务备注。当memo_mode为1时，memo将采用追加方式。
	_memo string
	// 状态值，多个以逗号分隔
	_status string
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

// NewTaobaoQianniuTaskUpdateRequest 初始化TaobaoQianniuTaskUpdateAPIRequest对象
func NewTaobaoQianniuTaskUpdateRequest() *TaobaoQianniuTaskUpdateAPIRequest {
	return &TaobaoQianniuTaskUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQianniuTaskUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.qianniu.task.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQianniuTaskUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSubStatus is SubStatus Setter
// 子任务状态，由业务方自定义
func (r *TaobaoQianniuTaskUpdateAPIRequest) SetSubStatus(_subStatus string) error {
	r._subStatus = _subStatus
	r.Set("sub_status", _subStatus)
	return nil
}

// GetSubStatus SubStatus Getter
func (r TaobaoQianniuTaskUpdateAPIRequest) GetSubStatus() string {
	return r._subStatus
}

// SetTag is Tag Setter
// 任务标签
func (r *TaobaoQianniuTaskUpdateAPIRequest) SetTag(_tag string) error {
	r._tag = _tag
	r.Set("tag", _tag)
	return nil
}

// GetTag Tag Getter
func (r TaobaoQianniuTaskUpdateAPIRequest) GetTag() string {
	return r._tag
}

// SetMemo is Memo Setter
// 任务备注。当memo_mode为1时，memo将采用追加方式。
func (r *TaobaoQianniuTaskUpdateAPIRequest) SetMemo(_memo string) error {
	r._memo = _memo
	r.Set("memo", _memo)
	return nil
}

// GetMemo Memo Getter
func (r TaobaoQianniuTaskUpdateAPIRequest) GetMemo() string {
	return r._memo
}

// SetStatus is Status Setter
// 状态值，多个以逗号分隔
func (r *TaobaoQianniuTaskUpdateAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoQianniuTaskUpdateAPIRequest) GetStatus() string {
	return r._status
}

// SetBizParam is BizParam Setter
// 应用自定义参数
func (r *TaobaoQianniuTaskUpdateAPIRequest) SetBizParam(_bizParam string) error {
	r._bizParam = _bizParam
	r.Set("biz_param", _bizParam)
	return nil
}

// GetBizParam BizParam Getter
func (r TaobaoQianniuTaskUpdateAPIRequest) GetBizParam() string {
	return r._bizParam
}

// SetTaskId is TaskId Setter
// 任务ID
func (r *TaobaoQianniuTaskUpdateAPIRequest) SetTaskId(_taskId int64) error {
	r._taskId = _taskId
	r.Set("task_id", _taskId)
	return nil
}

// GetTaskId TaskId Getter
func (r TaobaoQianniuTaskUpdateAPIRequest) GetTaskId() int64 {
	return r._taskId
}

// SetRemindTime is RemindTime Setter
// 提醒时间，时间的毫秒数
func (r *TaobaoQianniuTaskUpdateAPIRequest) SetRemindTime(_remindTime int64) error {
	r._remindTime = _remindTime
	r.Set("remind_time", _remindTime)
	return nil
}

// GetRemindTime RemindTime Getter
func (r TaobaoQianniuTaskUpdateAPIRequest) GetRemindTime() int64 {
	return r._remindTime
}

// SetRemindFlag is RemindFlag Setter
// 0为不提醒，1为全部提醒，2为PC提醒，3为移动提醒，4为已提醒，5为已忽略。
func (r *TaobaoQianniuTaskUpdateAPIRequest) SetRemindFlag(_remindFlag int64) error {
	r._remindFlag = _remindFlag
	r.Set("remind_flag", _remindFlag)
	return nil
}

// GetRemindFlag RemindFlag Getter
func (r TaobaoQianniuTaskUpdateAPIRequest) GetRemindFlag() int64 {
	return r._remindFlag
}

// SetMemoMode is MemoMode Setter
// 表示memo字段的更新策略。如需采用追加方式的，请将此字段设置为1。
func (r *TaobaoQianniuTaskUpdateAPIRequest) SetMemoMode(_memoMode int64) error {
	r._memoMode = _memoMode
	r.Set("memo_mode", _memoMode)
	return nil
}

// GetMemoMode MemoMode Getter
func (r TaobaoQianniuTaskUpdateAPIRequest) GetMemoMode() int64 {
	return r._memoMode
}

// SetPriority is Priority Setter
// 默认填0，数字越大优化级越高。当前常用0和1.
func (r *TaobaoQianniuTaskUpdateAPIRequest) SetPriority(_priority int64) error {
	r._priority = _priority
	r.Set("priority", _priority)
	return nil
}

// GetPriority Priority Getter
func (r TaobaoQianniuTaskUpdateAPIRequest) GetPriority() int64 {
	return r._priority
}

// SetIsDeleted is IsDeleted Setter
// 0表示没有删除，1表示删除
func (r *TaobaoQianniuTaskUpdateAPIRequest) SetIsDeleted(_isDeleted int64) error {
	r._isDeleted = _isDeleted
	r.Set("is_deleted", _isDeleted)
	return nil
}

// GetIsDeleted IsDeleted Getter
func (r TaobaoQianniuTaskUpdateAPIRequest) GetIsDeleted() int64 {
	return r._isDeleted
}
