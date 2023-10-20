package iotticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoiotticketspmaintainvtwocreateAPIRequest 服务商制定维修费方案 API请求
// cainiao.iot.ticket.sp.maintain.vtwo.create
//
// 服务商制定维修费方案
type CainiaoiotticketspmaintainvtwocreateAPIRequest struct {
	model.Params
	// 维修方案
	_makeMaintainPlanTopRequest *MakeMaintainPlanV2topRequest
}

// NewCainiaoiotticketspmaintainvtwocreateRequest 初始化CainiaoiotticketspmaintainvtwocreateAPIRequest对象
func NewCainiaoiotticketspmaintainvtwocreateRequest() *CainiaoiotticketspmaintainvtwocreateAPIRequest {
	return &CainiaoiotticketspmaintainvtwocreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoiotticketspmaintainvtwocreateAPIRequest) GetApiMethodName() string {
	return "cainiao.iot.ticket.sp.maintain.vtwo.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoiotticketspmaintainvtwocreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoiotticketspmaintainvtwocreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMakeMaintainPlanTopRequest is MakeMaintainPlanTopRequest Setter
// 维修方案
func (r *CainiaoiotticketspmaintainvtwocreateAPIRequest) SetMakeMaintainPlanTopRequest(_makeMaintainPlanTopRequest *MakeMaintainPlanV2topRequest) error {
	r._makeMaintainPlanTopRequest = _makeMaintainPlanTopRequest
	r.Set("make_maintain_plan_top_request", _makeMaintainPlanTopRequest)
	return nil
}

// GetMakeMaintainPlanTopRequest MakeMaintainPlanTopRequest Getter
func (r CainiaoiotticketspmaintainvtwocreateAPIRequest) GetMakeMaintainPlanTopRequest() *MakeMaintainPlanV2topRequest {
	return r._makeMaintainPlanTopRequest
}
