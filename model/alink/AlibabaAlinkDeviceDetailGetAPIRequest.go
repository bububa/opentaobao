package alink

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalinkdevicedetailgetAPIRequest 获取设备详情 API请求
// alibaba.alink.device.detail.get
//
// 阿里智能获取设备详情
type AlibabaalinkdevicedetailgetAPIRequest struct {
	model.Params
	// 设备id
	_uuid string
}

// NewAlibabaalinkdevicedetailgetRequest 初始化AlibabaalinkdevicedetailgetAPIRequest对象
func NewAlibabaalinkdevicedetailgetRequest() *AlibabaalinkdevicedetailgetAPIRequest {
	return &AlibabaalinkdevicedetailgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalinkdevicedetailgetAPIRequest) GetApiMethodName() string {
	return "alibaba.alink.device.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalinkdevicedetailgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalinkdevicedetailgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// 设备id
func (r *AlibabaalinkdevicedetailgetAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaalinkdevicedetailgetAPIRequest) GetUuid() string {
	return r._uuid
}
