package alink

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalinkdeviceunifystatussetAPIRequest 设置设备标准属性状态 API请求
// alibaba.alink.device.unify.status.set
//
// 操作用户绑定的设备
type AlibabaalinkdeviceunifystatussetAPIRequest struct {
	model.Params
	// uuid
	_uuid string
	// 设备的设置参数数据
	_instructions string
}

// NewAlibabaalinkdeviceunifystatussetRequest 初始化AlibabaalinkdeviceunifystatussetAPIRequest对象
func NewAlibabaalinkdeviceunifystatussetRequest() *AlibabaalinkdeviceunifystatussetAPIRequest {
	return &AlibabaalinkdeviceunifystatussetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalinkdeviceunifystatussetAPIRequest) GetApiMethodName() string {
	return "alibaba.alink.device.unify.status.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalinkdeviceunifystatussetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalinkdeviceunifystatussetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// uuid
func (r *AlibabaalinkdeviceunifystatussetAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaalinkdeviceunifystatussetAPIRequest) GetUuid() string {
	return r._uuid
}

// SetInstructions is Instructions Setter
// 设备的设置参数数据
func (r *AlibabaalinkdeviceunifystatussetAPIRequest) SetInstructions(_instructions string) error {
	r._instructions = _instructions
	r.Set("instructions", _instructions)
	return nil
}

// GetInstructions Instructions Getter
func (r AlibabaalinkdeviceunifystatussetAPIRequest) GetInstructions() string {
	return r._instructions
}
