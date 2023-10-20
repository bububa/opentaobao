package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusdeviceopenapifeedbackeventinfoAPIRequest IVS事件处理反馈接口 API请求
// alibaba.campus.device.openapi.feedbackeventinfo
//
// 提供给第三方ISV的的事件信息处理反馈的接口
type AlibabacampusdeviceopenapifeedbackeventinfoAPIRequest struct {
	model.Params
	// 系统上下文
	_param0 *WorkBenchContext
	// 请求封装类
	_param1 *EventInfoApiDto
}

// NewAlibabacampusdeviceopenapifeedbackeventinfoRequest 初始化AlibabacampusdeviceopenapifeedbackeventinfoAPIRequest对象
func NewAlibabacampusdeviceopenapifeedbackeventinfoRequest() *AlibabacampusdeviceopenapifeedbackeventinfoAPIRequest {
	return &AlibabacampusdeviceopenapifeedbackeventinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusdeviceopenapifeedbackeventinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.device.openapi.feedbackeventinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusdeviceopenapifeedbackeventinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusdeviceopenapifeedbackeventinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 系统上下文
func (r *AlibabacampusdeviceopenapifeedbackeventinfoAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabacampusdeviceopenapifeedbackeventinfoAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// SetParam1 is Param1 Setter
// 请求封装类
func (r *AlibabacampusdeviceopenapifeedbackeventinfoAPIRequest) SetParam1(_param1 *EventInfoApiDto) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabacampusdeviceopenapifeedbackeventinfoAPIRequest) GetParam1() *EventInfoApiDto {
	return r._param1
}
