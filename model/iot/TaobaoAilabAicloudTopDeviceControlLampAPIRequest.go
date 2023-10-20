package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopdevicecontrollampAPIRequest 台灯控制 API请求
// taobao.ailab.aicloud.top.device.control.lamp
//
// 台灯控制
type TaobaoailabaicloudtopdevicecontrollampAPIRequest struct {
	model.Params
	// 设备id
	_param1 string
	// 目标名称
	_param3 string
	// 用户信息
	_param0 *OpenBaseInfo
	// 是否打开
	_param2 bool
}

// NewTaobaoailabaicloudtopdevicecontrollampRequest 初始化TaobaoailabaicloudtopdevicecontrollampAPIRequest对象
func NewTaobaoailabaicloudtopdevicecontrollampRequest() *TaobaoailabaicloudtopdevicecontrollampAPIRequest {
	return &TaobaoailabaicloudtopdevicecontrollampAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtopdevicecontrollampAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.device.control.lamp"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtopdevicecontrollampAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtopdevicecontrollampAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam1 is Param1 Setter
// 设备id
func (r *TaobaoailabaicloudtopdevicecontrollampAPIRequest) SetParam1(_param1 string) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r TaobaoailabaicloudtopdevicecontrollampAPIRequest) GetParam1() string {
	return r._param1
}

// SetParam3 is Param3 Setter
// 目标名称
func (r *TaobaoailabaicloudtopdevicecontrollampAPIRequest) SetParam3(_param3 string) error {
	r._param3 = _param3
	r.Set("param3", _param3)
	return nil
}

// GetParam3 Param3 Getter
func (r TaobaoailabaicloudtopdevicecontrollampAPIRequest) GetParam3() string {
	return r._param3
}

// SetParam0 is Param0 Setter
// 用户信息
func (r *TaobaoailabaicloudtopdevicecontrollampAPIRequest) SetParam0(_param0 *OpenBaseInfo) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoailabaicloudtopdevicecontrollampAPIRequest) GetParam0() *OpenBaseInfo {
	return r._param0
}

// SetParam2 is Param2 Setter
// 是否打开
func (r *TaobaoailabaicloudtopdevicecontrollampAPIRequest) SetParam2(_param2 bool) error {
	r._param2 = _param2
	r.Set("param2", _param2)
	return nil
}

// GetParam2 Param2 Getter
func (r TaobaoailabaicloudtopdevicecontrollampAPIRequest) GetParam2() bool {
	return r._param2
}
