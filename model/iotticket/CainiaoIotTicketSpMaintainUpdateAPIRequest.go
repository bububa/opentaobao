package iotticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoiotticketspmaintainupdateAPIRequest IoT售后服务商维修方案更新 API请求
// cainiao.iot.ticket.sp.maintain.update
//
// IoT售后服务商维修方案更新
type CainiaoiotticketspmaintainupdateAPIRequest struct {
	model.Params
	// 请求参数
	_param *UpdateMaintainPlanTopRequest
}

// NewCainiaoiotticketspmaintainupdateRequest 初始化CainiaoiotticketspmaintainupdateAPIRequest对象
func NewCainiaoiotticketspmaintainupdateRequest() *CainiaoiotticketspmaintainupdateAPIRequest {
	return &CainiaoiotticketspmaintainupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoiotticketspmaintainupdateAPIRequest) GetApiMethodName() string {
	return "cainiao.iot.ticket.sp.maintain.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoiotticketspmaintainupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoiotticketspmaintainupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 请求参数
func (r *CainiaoiotticketspmaintainupdateAPIRequest) SetParam(_param *UpdateMaintainPlanTopRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r CainiaoiotticketspmaintainupdateAPIRequest) GetParam() *UpdateMaintainPlanTopRequest {
	return r._param
}
