package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopdevicecontrolcustomAPIRequest 设备控制自定义扩展接口 API请求
// taobao.ailab.aicloud.top.device.control.custom
//
// 设备控制自定义扩展接口
type TaobaoailabaicloudtopdevicecontrolcustomAPIRequest struct {
	model.Params
	// 参数key-value列表
	_param2 []string
	// 设备id
	_param1 string
	// 用户信息
	_param0 *OpenBaseInfo
}

// NewTaobaoailabaicloudtopdevicecontrolcustomRequest 初始化TaobaoailabaicloudtopdevicecontrolcustomAPIRequest对象
func NewTaobaoailabaicloudtopdevicecontrolcustomRequest() *TaobaoailabaicloudtopdevicecontrolcustomAPIRequest {
	return &TaobaoailabaicloudtopdevicecontrolcustomAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtopdevicecontrolcustomAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.device.control.custom"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtopdevicecontrolcustomAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtopdevicecontrolcustomAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam2 is Param2 Setter
// 参数key-value列表
func (r *TaobaoailabaicloudtopdevicecontrolcustomAPIRequest) SetParam2(_param2 []string) error {
	r._param2 = _param2
	r.Set("param2", _param2)
	return nil
}

// GetParam2 Param2 Getter
func (r TaobaoailabaicloudtopdevicecontrolcustomAPIRequest) GetParam2() []string {
	return r._param2
}

// SetParam1 is Param1 Setter
// 设备id
func (r *TaobaoailabaicloudtopdevicecontrolcustomAPIRequest) SetParam1(_param1 string) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r TaobaoailabaicloudtopdevicecontrolcustomAPIRequest) GetParam1() string {
	return r._param1
}

// SetParam0 is Param0 Setter
// 用户信息
func (r *TaobaoailabaicloudtopdevicecontrolcustomAPIRequest) SetParam0(_param0 *OpenBaseInfo) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoailabaicloudtopdevicecontrolcustomAPIRequest) GetParam0() *OpenBaseInfo {
	return r._param0
}
