package iot

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest 点播url API请求
// taobao.ailab.aicloud.top.device.control.playurl
//
// 点播url
type TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest struct {
	model.Params
	// 设备id
	_param1 string
	// url
	_param2 string
	// 用户信息
	_param0 *OpenBaseInfo
}

// NewTaobaoAilabAicloudTopDeviceControlPlayurlRequest 初始化TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest对象
func NewTaobaoAilabAicloudTopDeviceControlPlayurlRequest() *TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest {
	return &TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest) Reset() {
	r._param1 = ""
	r._param2 = ""
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.device.control.playurl"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam1 is Param1 Setter
// 设备id
func (r *TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest) SetParam1(_param1 string) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest) GetParam1() string {
	return r._param1
}

// SetParam2 is Param2 Setter
// url
func (r *TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest) SetParam2(_param2 string) error {
	r._param2 = _param2
	r.Set("param2", _param2)
	return nil
}

// GetParam2 Param2 Getter
func (r TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest) GetParam2() string {
	return r._param2
}

// SetParam0 is Param0 Setter
// 用户信息
func (r *TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest) SetParam0(_param0 *OpenBaseInfo) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest) GetParam0() *OpenBaseInfo {
	return r._param0
}

var poolTaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAilabAicloudTopDeviceControlPlayurlRequest()
	},
}

// GetTaobaoAilabAicloudTopDeviceControlPlayurlRequest 从 sync.Pool 获取 TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest
func GetTaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest() *TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest {
	return poolTaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest.Get().(*TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest)
}

// ReleaseTaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest 将 TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest 放入 sync.Pool
func ReleaseTaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest(v *TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest) {
	v.Reset()
	poolTaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest.Put(v)
}
