package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopdevicegetstatusAPIRequest 获取设备状态 API请求
// taobao.ailab.aicloud.top.device.getstatus
//
// 获取设备状态
type TaobaoailabaicloudtopdevicegetstatusAPIRequest struct {
	model.Params
	// 设备id
	_param1 string
	// 用户信息
	_param0 *OpenBaseInfo
}

// NewTaobaoailabaicloudtopdevicegetstatusRequest 初始化TaobaoailabaicloudtopdevicegetstatusAPIRequest对象
func NewTaobaoailabaicloudtopdevicegetstatusRequest() *TaobaoailabaicloudtopdevicegetstatusAPIRequest {
	return &TaobaoailabaicloudtopdevicegetstatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtopdevicegetstatusAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.device.getstatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtopdevicegetstatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtopdevicegetstatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam1 is Param1 Setter
// 设备id
func (r *TaobaoailabaicloudtopdevicegetstatusAPIRequest) SetParam1(_param1 string) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r TaobaoailabaicloudtopdevicegetstatusAPIRequest) GetParam1() string {
	return r._param1
}

// SetParam0 is Param0 Setter
// 用户信息
func (r *TaobaoailabaicloudtopdevicegetstatusAPIRequest) SetParam0(_param0 *OpenBaseInfo) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoailabaicloudtopdevicegetstatusAPIRequest) GetParam0() *OpenBaseInfo {
	return r._param0
}
