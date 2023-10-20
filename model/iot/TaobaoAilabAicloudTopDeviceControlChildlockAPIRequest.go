package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopdevicecontrolchildlockAPIRequest 设备儿童锁 API请求
// taobao.ailab.aicloud.top.device.control.childlock
//
// 设备儿童锁
type TaobaoailabaicloudtopdevicecontrolchildlockAPIRequest struct {
	model.Params
	// 设备id
	_param1 string
	// 用户信息
	_param0 *OpenBaseInfo
	// 是否打开
	_param2 bool
}

// NewTaobaoailabaicloudtopdevicecontrolchildlockRequest 初始化TaobaoailabaicloudtopdevicecontrolchildlockAPIRequest对象
func NewTaobaoailabaicloudtopdevicecontrolchildlockRequest() *TaobaoailabaicloudtopdevicecontrolchildlockAPIRequest {
	return &TaobaoailabaicloudtopdevicecontrolchildlockAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtopdevicecontrolchildlockAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.device.control.childlock"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtopdevicecontrolchildlockAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtopdevicecontrolchildlockAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam1 is Param1 Setter
// 设备id
func (r *TaobaoailabaicloudtopdevicecontrolchildlockAPIRequest) SetParam1(_param1 string) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r TaobaoailabaicloudtopdevicecontrolchildlockAPIRequest) GetParam1() string {
	return r._param1
}

// SetParam0 is Param0 Setter
// 用户信息
func (r *TaobaoailabaicloudtopdevicecontrolchildlockAPIRequest) SetParam0(_param0 *OpenBaseInfo) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoailabaicloudtopdevicecontrolchildlockAPIRequest) GetParam0() *OpenBaseInfo {
	return r._param0
}

// SetParam2 is Param2 Setter
// 是否打开
func (r *TaobaoailabaicloudtopdevicecontrolchildlockAPIRequest) SetParam2(_param2 bool) error {
	r._param2 = _param2
	r.Set("param2", _param2)
	return nil
}

// GetParam2 Param2 Getter
func (r TaobaoailabaicloudtopdevicecontrolchildlockAPIRequest) GetParam2() bool {
	return r._param2
}
