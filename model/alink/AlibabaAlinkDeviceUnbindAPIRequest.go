package alink

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalinkdeviceunbindAPIRequest 解绑设备 API请求
// alibaba.alink.device.unbind
//
// 阿里智能解绑设备
type AlibabaalinkdeviceunbindAPIRequest struct {
	model.Params
	// 设备id
	_uuid string
}

// NewAlibabaalinkdeviceunbindRequest 初始化AlibabaalinkdeviceunbindAPIRequest对象
func NewAlibabaalinkdeviceunbindRequest() *AlibabaalinkdeviceunbindAPIRequest {
	return &AlibabaalinkdeviceunbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalinkdeviceunbindAPIRequest) GetApiMethodName() string {
	return "alibaba.alink.device.unbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalinkdeviceunbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalinkdeviceunbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// 设备id
func (r *AlibabaalinkdeviceunbindAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaalinkdeviceunbindAPIRequest) GetUuid() string {
	return r._uuid
}
