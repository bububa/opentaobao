package alink

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalinkdeviceunifystatusgetAPIRequest 查询设备标准属性最新状态 API请求
// alibaba.alink.device.unify.status.get
//
// 查询设备最新标准属性状态
type AlibabaalinkdeviceunifystatusgetAPIRequest struct {
	model.Params
	// 设备id
	_uuid string
}

// NewAlibabaalinkdeviceunifystatusgetRequest 初始化AlibabaalinkdeviceunifystatusgetAPIRequest对象
func NewAlibabaalinkdeviceunifystatusgetRequest() *AlibabaalinkdeviceunifystatusgetAPIRequest {
	return &AlibabaalinkdeviceunifystatusgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalinkdeviceunifystatusgetAPIRequest) GetApiMethodName() string {
	return "alibaba.alink.device.unify.status.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalinkdeviceunifystatusgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalinkdeviceunifystatusgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// 设备id
func (r *AlibabaalinkdeviceunifystatusgetAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaalinkdeviceunifystatusgetAPIRequest) GetUuid() string {
	return r._uuid
}
