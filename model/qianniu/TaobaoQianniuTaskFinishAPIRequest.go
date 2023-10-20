package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqianniutaskfinishAPIRequest 完成轻任务 API请求
// taobao.qianniu.task.finish
//
// 由任务执行者调用
type TaobaoqianniutaskfinishAPIRequest struct {
	model.Params
	// 任务备注
	_memo string
	// 任务ID
	_taskId int64
}

// NewTaobaoqianniutaskfinishRequest 初始化TaobaoqianniutaskfinishAPIRequest对象
func NewTaobaoqianniutaskfinishRequest() *TaobaoqianniutaskfinishAPIRequest {
	return &TaobaoqianniutaskfinishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqianniutaskfinishAPIRequest) GetApiMethodName() string {
	return "taobao.qianniu.task.finish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqianniutaskfinishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqianniutaskfinishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMemo is Memo Setter
// 任务备注
func (r *TaobaoqianniutaskfinishAPIRequest) SetMemo(_memo string) error {
	r._memo = _memo
	r.Set("memo", _memo)
	return nil
}

// GetMemo Memo Getter
func (r TaobaoqianniutaskfinishAPIRequest) GetMemo() string {
	return r._memo
}

// SetTaskId is TaskId Setter
// 任务ID
func (r *TaobaoqianniutaskfinishAPIRequest) SetTaskId(_taskId int64) error {
	r._taskId = _taskId
	r.Set("task_id", _taskId)
	return nil
}

// GetTaskId TaskId Getter
func (r TaobaoqianniutaskfinishAPIRequest) GetTaskId() int64 {
	return r._taskId
}
