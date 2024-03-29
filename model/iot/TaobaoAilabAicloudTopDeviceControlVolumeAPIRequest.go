package iot

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest 设备音量 API请求
// taobao.ailab.aicloud.top.device.control.volume
//
// 设备音量
type TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest struct {
	model.Params
	// 设备id
	_param1 string
	// 用户信息
	_param0 *OpenBaseInfo
	// 音量0-100
	_param2 int64
}

// NewTaobaoAilabAicloudTopDeviceControlVolumeRequest 初始化TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest对象
func NewTaobaoAilabAicloudTopDeviceControlVolumeRequest() *TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest {
	return &TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest) Reset() {
	r._param1 = ""
	r._param0 = nil
	r._param2 = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.device.control.volume"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam1 is Param1 Setter
// 设备id
func (r *TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest) SetParam1(_param1 string) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest) GetParam1() string {
	return r._param1
}

// SetParam0 is Param0 Setter
// 用户信息
func (r *TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest) SetParam0(_param0 *OpenBaseInfo) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest) GetParam0() *OpenBaseInfo {
	return r._param0
}

// SetParam2 is Param2 Setter
// 音量0-100
func (r *TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest) SetParam2(_param2 int64) error {
	r._param2 = _param2
	r.Set("param2", _param2)
	return nil
}

// GetParam2 Param2 Getter
func (r TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest) GetParam2() int64 {
	return r._param2
}

var poolTaobaoAilabAicloudTopDeviceControlVolumeAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAilabAicloudTopDeviceControlVolumeRequest()
	},
}

// GetTaobaoAilabAicloudTopDeviceControlVolumeRequest 从 sync.Pool 获取 TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest
func GetTaobaoAilabAicloudTopDeviceControlVolumeAPIRequest() *TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest {
	return poolTaobaoAilabAicloudTopDeviceControlVolumeAPIRequest.Get().(*TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest)
}

// ReleaseTaobaoAilabAicloudTopDeviceControlVolumeAPIRequest 将 TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest 放入 sync.Pool
func ReleaseTaobaoAilabAicloudTopDeviceControlVolumeAPIRequest(v *TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest) {
	v.Reset()
	poolTaobaoAilabAicloudTopDeviceControlVolumeAPIRequest.Put(v)
}
