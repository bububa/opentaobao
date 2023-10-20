package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusdeviceopenapisaveeventinfoforibosAPIRequest saveeventinfoforibos API请求
// alibaba.campus.device.openapi.saveeventinfoforibos
//
// IBos的事件信息上报与反馈的处理接口
type AlibabacampusdeviceopenapisaveeventinfoforibosAPIRequest struct {
	model.Params
	// 系统自动生成
	_param0 *WorkBenchContext
	// 系统自动生成
	_param1 *EventInfoApiDto
}

// NewAlibabacampusdeviceopenapisaveeventinfoforibosRequest 初始化AlibabacampusdeviceopenapisaveeventinfoforibosAPIRequest对象
func NewAlibabacampusdeviceopenapisaveeventinfoforibosRequest() *AlibabacampusdeviceopenapisaveeventinfoforibosAPIRequest {
	return &AlibabacampusdeviceopenapisaveeventinfoforibosAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusdeviceopenapisaveeventinfoforibosAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.device.openapi.saveeventinfoforibos"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusdeviceopenapisaveeventinfoforibosAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusdeviceopenapisaveeventinfoforibosAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 系统自动生成
func (r *AlibabacampusdeviceopenapisaveeventinfoforibosAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabacampusdeviceopenapisaveeventinfoforibosAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// SetParam1 is Param1 Setter
// 系统自动生成
func (r *AlibabacampusdeviceopenapisaveeventinfoforibosAPIRequest) SetParam1(_param1 *EventInfoApiDto) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabacampusdeviceopenapisaveeventinfoforibosAPIRequest) GetParam1() *EventInfoApiDto {
	return r._param1
}
