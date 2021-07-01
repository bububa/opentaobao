package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest
设备音量 API请求
taobao.ailab.aicloud.top.device.control.volume

设备音量 */
type TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest struct {
	model.Params
	// 用户信息
	_param0 *OpenBaseInfo
	// 设备id
	_param1 string
	// 音量0-100
	_param2 int64
}

// NewTaobaoAilabAicloudTopDeviceControlVolumeRequest 初始化TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest对象
func NewTaobaoAilabAicloudTopDeviceControlVolumeRequest() *TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest {
	return &TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.device.control.volume"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 用户信息
func (r *TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest) SetParam0(_param0 *OpenBaseInfo) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest) GetParam0() *OpenBaseInfo {
	return r._param0
}

// Set is Param1 Setter
// 设备id
func (r *TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest) SetParam1(_param1 string) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// Get Param1 Getter
func (r TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest) GetParam1() string {
	return r._param1
}

// Set is Param2 Setter
// 音量0-100
func (r *TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest) SetParam2(_param2 int64) error {
	r._param2 = _param2
	r.Set("param2", _param2)
	return nil
}

// Get Param2 Getter
func (r TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest) GetParam2() int64 {
	return r._param2
}
