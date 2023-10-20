package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabajymindustryinformationcallbakAPIRequest VMOS回调行业信息系统 API请求
// alibaba.jym.industry.information.callbak
//
// VMOS回调交易猫行业信息系统
type AlibabajymindustryinformationcallbakAPIRequest struct {
	model.Params
	// 任务ID
	_taskId string
	// 幂等ID
	_bizId string
	// 内容
	_content string
	// 状态
	_status int64
}

// NewAlibabajymindustryinformationcallbakRequest 初始化AlibabajymindustryinformationcallbakAPIRequest对象
func NewAlibabajymindustryinformationcallbakRequest() *AlibabajymindustryinformationcallbakAPIRequest {
	return &AlibabajymindustryinformationcallbakAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabajymindustryinformationcallbakAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.industry.information.callbak"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabajymindustryinformationcallbakAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabajymindustryinformationcallbakAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTaskId is TaskId Setter
// 任务ID
func (r *AlibabajymindustryinformationcallbakAPIRequest) SetTaskId(_taskId string) error {
	r._taskId = _taskId
	r.Set("task_id", _taskId)
	return nil
}

// GetTaskId TaskId Getter
func (r AlibabajymindustryinformationcallbakAPIRequest) GetTaskId() string {
	return r._taskId
}

// SetBizId is BizId Setter
// 幂等ID
func (r *AlibabajymindustryinformationcallbakAPIRequest) SetBizId(_bizId string) error {
	r._bizId = _bizId
	r.Set("biz_id", _bizId)
	return nil
}

// GetBizId BizId Getter
func (r AlibabajymindustryinformationcallbakAPIRequest) GetBizId() string {
	return r._bizId
}

// SetContent is Content Setter
// 内容
func (r *AlibabajymindustryinformationcallbakAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r AlibabajymindustryinformationcallbakAPIRequest) GetContent() string {
	return r._content
}

// SetStatus is Status Setter
// 状态
func (r *AlibabajymindustryinformationcallbakAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabajymindustryinformationcallbakAPIRequest) GetStatus() int64 {
	return r._status
}
