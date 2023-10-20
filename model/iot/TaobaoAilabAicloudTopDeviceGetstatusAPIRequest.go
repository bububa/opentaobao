package iot

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceGetstatusAPIRequest 获取设备状态 API请求
// taobao.ailab.aicloud.top.device.getstatus
//
// 获取设备状态
type TaobaoAilabAicloudTopDeviceGetstatusAPIRequest struct {
	model.Params
	// 设备id
	_param1 string
	// 用户信息
	_param0 *OpenBaseInfo
}

// NewTaobaoAilabAicloudTopDeviceGetstatusRequest 初始化TaobaoAilabAicloudTopDeviceGetstatusAPIRequest对象
func NewTaobaoAilabAicloudTopDeviceGetstatusRequest() *TaobaoAilabAicloudTopDeviceGetstatusAPIRequest {
	return &TaobaoAilabAicloudTopDeviceGetstatusAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAilabAicloudTopDeviceGetstatusAPIRequest) Reset() {
	r._param1 = ""
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceGetstatusAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.device.getstatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceGetstatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAilabAicloudTopDeviceGetstatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam1 is Param1 Setter
// 设备id
func (r *TaobaoAilabAicloudTopDeviceGetstatusAPIRequest) SetParam1(_param1 string) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r TaobaoAilabAicloudTopDeviceGetstatusAPIRequest) GetParam1() string {
	return r._param1
}

// SetParam0 is Param0 Setter
// 用户信息
func (r *TaobaoAilabAicloudTopDeviceGetstatusAPIRequest) SetParam0(_param0 *OpenBaseInfo) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoAilabAicloudTopDeviceGetstatusAPIRequest) GetParam0() *OpenBaseInfo {
	return r._param0
}

var poolTaobaoAilabAicloudTopDeviceGetstatusAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAilabAicloudTopDeviceGetstatusRequest()
	},
}

// GetTaobaoAilabAicloudTopDeviceGetstatusRequest 从 sync.Pool 获取 TaobaoAilabAicloudTopDeviceGetstatusAPIRequest
func GetTaobaoAilabAicloudTopDeviceGetstatusAPIRequest() *TaobaoAilabAicloudTopDeviceGetstatusAPIRequest {
	return poolTaobaoAilabAicloudTopDeviceGetstatusAPIRequest.Get().(*TaobaoAilabAicloudTopDeviceGetstatusAPIRequest)
}

// ReleaseTaobaoAilabAicloudTopDeviceGetstatusAPIRequest 将 TaobaoAilabAicloudTopDeviceGetstatusAPIRequest 放入 sync.Pool
func ReleaseTaobaoAilabAicloudTopDeviceGetstatusAPIRequest(v *TaobaoAilabAicloudTopDeviceGetstatusAPIRequest) {
	v.Reset()
	poolTaobaoAilabAicloudTopDeviceGetstatusAPIRequest.Put(v)
}
