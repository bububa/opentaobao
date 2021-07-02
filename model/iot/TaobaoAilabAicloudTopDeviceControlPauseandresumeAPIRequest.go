package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest 设备播放暂停 API请求
// taobao.ailab.aicloud.top.device.control.pauseandresume
//
// 设备播放暂停
type TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest struct {
	model.Params
	// 用户信息
	_param0 *OpenBaseInfo
	// 设备id
	_param1 string
	// 是暂停还是继续
	_param2 bool
}

// NewTaobaoAilabAicloudTopDeviceControlPauseandresumeRequest 初始化TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest对象
func NewTaobaoAilabAicloudTopDeviceControlPauseandresumeRequest() *TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest {
	return &TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.device.control.pauseandresume"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 用户信息
func (r *TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest) SetParam0(_param0 *OpenBaseInfo) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest) GetParam0() *OpenBaseInfo {
	return r._param0
}

// Set is Param1 Setter
// 设备id
func (r *TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest) SetParam1(_param1 string) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// Get Param1 Getter
func (r TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest) GetParam1() string {
	return r._param1
}

// Set is Param2 Setter
// 是暂停还是继续
func (r *TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest) SetParam2(_param2 bool) error {
	r._param2 = _param2
	r.Set("param2", _param2)
	return nil
}

// Get Param2 Getter
func (r TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest) GetParam2() bool {
	return r._param2
}
