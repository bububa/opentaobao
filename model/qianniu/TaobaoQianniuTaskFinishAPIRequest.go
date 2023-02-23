package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQianniuTaskFinishAPIRequest 完成轻任务 API请求
// taobao.qianniu.task.finish
//
// 由任务执行者调用
type TaobaoQianniuTaskFinishAPIRequest struct {
	model.Params
	// 任务备注
	_memo string
	// 任务ID
	_taskId int64
}

// NewTaobaoQianniuTaskFinishRequest 初始化TaobaoQianniuTaskFinishAPIRequest对象
func NewTaobaoQianniuTaskFinishRequest() *TaobaoQianniuTaskFinishAPIRequest {
	return &TaobaoQianniuTaskFinishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQianniuTaskFinishAPIRequest) GetApiMethodName() string {
	return "taobao.qianniu.task.finish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQianniuTaskFinishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQianniuTaskFinishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMemo is Memo Setter
// 任务备注
func (r *TaobaoQianniuTaskFinishAPIRequest) SetMemo(_memo string) error {
	r._memo = _memo
	r.Set("memo", _memo)
	return nil
}

// GetMemo Memo Getter
func (r TaobaoQianniuTaskFinishAPIRequest) GetMemo() string {
	return r._memo
}

// SetTaskId is TaskId Setter
// 任务ID
func (r *TaobaoQianniuTaskFinishAPIRequest) SetTaskId(_taskId int64) error {
	r._taskId = _taskId
	r.Set("task_id", _taskId)
	return nil
}

// GetTaskId TaskId Getter
func (r TaobaoQianniuTaskFinishAPIRequest) GetTaskId() int64 {
	return r._taskId
}
