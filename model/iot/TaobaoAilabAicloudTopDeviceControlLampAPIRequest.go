package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopDeviceControlLampAPIRequest
台灯控制 API请求
taobao.ailab.aicloud.top.device.control.lamp

台灯控制 */
type TaobaoAilabAicloudTopDeviceControlLampAPIRequest struct {
	model.Params
	// 用户信息
	_param0 *OpenBaseInfo
	// 设备id
	_param1 string
	// 是否打开
	_param2 bool
	// 目标名称
	_param3 string
}

// NewTaobaoAilabAicloudTopDeviceControlLampRequest 初始化TaobaoAilabAicloudTopDeviceControlLampAPIRequest对象
func NewTaobaoAilabAicloudTopDeviceControlLampRequest() *TaobaoAilabAicloudTopDeviceControlLampAPIRequest {
	return &TaobaoAilabAicloudTopDeviceControlLampAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceControlLampAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.device.control.lamp"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceControlLampAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 用户信息
func (r *TaobaoAilabAicloudTopDeviceControlLampAPIRequest) SetParam0(_param0 *OpenBaseInfo) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r TaobaoAilabAicloudTopDeviceControlLampAPIRequest) GetParam0() *OpenBaseInfo {
	return r._param0
}

// Set is Param1 Setter
// 设备id
func (r *TaobaoAilabAicloudTopDeviceControlLampAPIRequest) SetParam1(_param1 string) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// Get Param1 Getter
func (r TaobaoAilabAicloudTopDeviceControlLampAPIRequest) GetParam1() string {
	return r._param1
}

// Set is Param2 Setter
// 是否打开
func (r *TaobaoAilabAicloudTopDeviceControlLampAPIRequest) SetParam2(_param2 bool) error {
	r._param2 = _param2
	r.Set("param2", _param2)
	return nil
}

// Get Param2 Getter
func (r TaobaoAilabAicloudTopDeviceControlLampAPIRequest) GetParam2() bool {
	return r._param2
}

// Set is Param3 Setter
// 目标名称
func (r *TaobaoAilabAicloudTopDeviceControlLampAPIRequest) SetParam3(_param3 string) error {
	r._param3 = _param3
	r.Set("param3", _param3)
	return nil
}

// Get Param3 Getter
func (r TaobaoAilabAicloudTopDeviceControlLampAPIRequest) GetParam3() string {
	return r._param3
}
