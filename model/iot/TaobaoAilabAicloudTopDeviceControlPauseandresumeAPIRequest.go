package iot

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest 设备播放暂停 API请求
// taobao.ailab.aicloud.top.device.control.pauseandresume
//
// 设备播放暂停
type TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest struct {
	model.Params
	// 设备id
	_param1 string
	// 用户信息
	_param0 *OpenBaseInfo
	// 是暂停还是继续
	_param2 bool
}

// NewTaobaoAilabAicloudTopDeviceControlPauseandresumeRequest 初始化TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest对象
func NewTaobaoAilabAicloudTopDeviceControlPauseandresumeRequest() *TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest {
	return &TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest) Reset() {
	r._param1 = ""
	r._param0 = nil
	r._param2 = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.device.control.pauseandresume"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam1 is Param1 Setter
// 设备id
func (r *TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest) SetParam1(_param1 string) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest) GetParam1() string {
	return r._param1
}

// SetParam0 is Param0 Setter
// 用户信息
func (r *TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest) SetParam0(_param0 *OpenBaseInfo) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest) GetParam0() *OpenBaseInfo {
	return r._param0
}

// SetParam2 is Param2 Setter
// 是暂停还是继续
func (r *TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest) SetParam2(_param2 bool) error {
	r._param2 = _param2
	r.Set("param2", _param2)
	return nil
}

// GetParam2 Param2 Getter
func (r TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest) GetParam2() bool {
	return r._param2
}

var poolTaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAilabAicloudTopDeviceControlPauseandresumeRequest()
	},
}

// GetTaobaoAilabAicloudTopDeviceControlPauseandresumeRequest 从 sync.Pool 获取 TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest
func GetTaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest() *TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest {
	return poolTaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest.Get().(*TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest)
}

// ReleaseTaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest 将 TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest 放入 sync.Pool
func ReleaseTaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest(v *TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest) {
	v.Reset()
	poolTaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest.Put(v)
}
