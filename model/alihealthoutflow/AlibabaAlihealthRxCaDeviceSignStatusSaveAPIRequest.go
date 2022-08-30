package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthRxCaDeviceSignStatusSaveAPIRequest 处方ca认证-厂商通知接口 API请求
// alibaba.alihealth.rx.ca.device.sign.status.save
//
// 处方ca认证-厂商通知接口
type AlibabaAlihealthRxCaDeviceSignStatusSaveAPIRequest struct {
	model.Params
	// 入参
	_param0 *CaSignStatusSaveRequest
}

// NewAlibabaAlihealthRxCaDeviceSignStatusSaveRequest 初始化AlibabaAlihealthRxCaDeviceSignStatusSaveAPIRequest对象
func NewAlibabaAlihealthRxCaDeviceSignStatusSaveRequest() *AlibabaAlihealthRxCaDeviceSignStatusSaveAPIRequest {
	return &AlibabaAlihealthRxCaDeviceSignStatusSaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthRxCaDeviceSignStatusSaveAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.rx.ca.device.sign.status.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthRxCaDeviceSignStatusSaveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam0 is Param0 Setter
// 入参
func (r *AlibabaAlihealthRxCaDeviceSignStatusSaveAPIRequest) SetParam0(_param0 *CaSignStatusSaveRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaAlihealthRxCaDeviceSignStatusSaveAPIRequest) GetParam0() *CaSignStatusSaveRequest {
	return r._param0
}
