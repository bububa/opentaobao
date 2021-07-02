package alink

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlinkDeviceDetailGetAPIRequest 获取设备详情 API请求
// alibaba.alink.device.detail.get
//
// 阿里智能获取设备详情
type AlibabaAlinkDeviceDetailGetAPIRequest struct {
	model.Params
	// 设备id
	_uuid string
}

// NewAlibabaAlinkDeviceDetailGetRequest 初始化AlibabaAlinkDeviceDetailGetAPIRequest对象
func NewAlibabaAlinkDeviceDetailGetRequest() *AlibabaAlinkDeviceDetailGetAPIRequest {
	return &AlibabaAlinkDeviceDetailGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlinkDeviceDetailGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alink.device.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlinkDeviceDetailGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetUuid is Uuid Setter
// 设备id
func (r *AlibabaAlinkDeviceDetailGetAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaAlinkDeviceDetailGetAPIRequest) GetUuid() string {
	return r._uuid
}
