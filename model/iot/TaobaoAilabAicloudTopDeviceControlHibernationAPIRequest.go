package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceControlHibernationAPIRequest 定时休眠 API请求
// taobao.ailab.aicloud.top.device.control.hibernation
//
// 定时休眠
type TaobaoAilabAicloudTopDeviceControlHibernationAPIRequest struct {
	model.Params
	// 设备id
	_param1 string
	// N秒后休眠
	_param2 string
	// 用户信息
	_param0 *OpenBaseInfo
}

// NewTaobaoAilabAicloudTopDeviceControlHibernationRequest 初始化TaobaoAilabAicloudTopDeviceControlHibernationAPIRequest对象
func NewTaobaoAilabAicloudTopDeviceControlHibernationRequest() *TaobaoAilabAicloudTopDeviceControlHibernationAPIRequest {
	return &TaobaoAilabAicloudTopDeviceControlHibernationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceControlHibernationAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.device.control.hibernation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceControlHibernationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAilabAicloudTopDeviceControlHibernationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam1 is Param1 Setter
// 设备id
func (r *TaobaoAilabAicloudTopDeviceControlHibernationAPIRequest) SetParam1(_param1 string) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r TaobaoAilabAicloudTopDeviceControlHibernationAPIRequest) GetParam1() string {
	return r._param1
}

// SetParam2 is Param2 Setter
// N秒后休眠
func (r *TaobaoAilabAicloudTopDeviceControlHibernationAPIRequest) SetParam2(_param2 string) error {
	r._param2 = _param2
	r.Set("param2", _param2)
	return nil
}

// GetParam2 Param2 Getter
func (r TaobaoAilabAicloudTopDeviceControlHibernationAPIRequest) GetParam2() string {
	return r._param2
}

// SetParam0 is Param0 Setter
// 用户信息
func (r *TaobaoAilabAicloudTopDeviceControlHibernationAPIRequest) SetParam0(_param0 *OpenBaseInfo) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoAilabAicloudTopDeviceControlHibernationAPIRequest) GetParam0() *OpenBaseInfo {
	return r._param0
}
