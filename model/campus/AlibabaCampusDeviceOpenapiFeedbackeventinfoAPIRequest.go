package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest IVS事件处理反馈接口 API请求
// alibaba.campus.device.openapi.feedbackeventinfo
//
// 提供给第三方ISV的的事件信息处理反馈的接口
type AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest struct {
	model.Params
	// 系统上下文
	_param0 *WorkBenchContext
	// 请求封装类
	_param1 *EventInfoApiDto
}

// NewAlibabaCampusDeviceOpenapiFeedbackeventinfoRequest 初始化AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest对象
func NewAlibabaCampusDeviceOpenapiFeedbackeventinfoRequest() *AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest {
	return &AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest) Reset() {
	r._param0 = nil
	r._param1 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.device.openapi.feedbackeventinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 系统上下文
func (r *AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// SetParam1 is Param1 Setter
// 请求封装类
func (r *AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest) SetParam1(_param1 *EventInfoApiDto) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest) GetParam1() *EventInfoApiDto {
	return r._param1
}

var poolAlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusDeviceOpenapiFeedbackeventinfoRequest()
	},
}

// GetAlibabaCampusDeviceOpenapiFeedbackeventinfoRequest 从 sync.Pool 获取 AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest
func GetAlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest() *AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest {
	return poolAlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest.Get().(*AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest)
}

// ReleaseAlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest 将 AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest(v *AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest) {
	v.Reset()
	poolAlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest.Put(v)
}
