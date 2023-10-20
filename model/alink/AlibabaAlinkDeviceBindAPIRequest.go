package alink

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalinkdevicebindAPIRequest 绑定设备 API请求
// alibaba.alink.device.bind
//
// 阿里智能解绑设备
type AlibabaalinkdevicebindAPIRequest struct {
	model.Params
	// 设备id
	_uuid string
}

// NewAlibabaalinkdevicebindRequest 初始化AlibabaalinkdevicebindAPIRequest对象
func NewAlibabaalinkdevicebindRequest() *AlibabaalinkdevicebindAPIRequest {
	return &AlibabaalinkdevicebindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalinkdevicebindAPIRequest) GetApiMethodName() string {
	return "alibaba.alink.device.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalinkdevicebindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalinkdevicebindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// 设备id
func (r *AlibabaalinkdevicebindAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaalinkdevicebindAPIRequest) GetUuid() string {
	return r._uuid
}
