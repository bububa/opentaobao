package iotticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoIotTicketSpMaintainVtwoCreateAPIRequest 服务商制定维修费方案 API请求
// cainiao.iot.ticket.sp.maintain.vtwo.create
//
// 服务商制定维修费方案
type CainiaoIotTicketSpMaintainVtwoCreateAPIRequest struct {
	model.Params
	// 维修方案
	_makeMaintainPlanTopRequest *MakeMaintainPlanV2TopRequest
}

// NewCainiaoIotTicketSpMaintainVtwoCreateRequest 初始化CainiaoIotTicketSpMaintainVtwoCreateAPIRequest对象
func NewCainiaoIotTicketSpMaintainVtwoCreateRequest() *CainiaoIotTicketSpMaintainVtwoCreateAPIRequest {
	return &CainiaoIotTicketSpMaintainVtwoCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoIotTicketSpMaintainVtwoCreateAPIRequest) GetApiMethodName() string {
	return "cainiao.iot.ticket.sp.maintain.vtwo.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoIotTicketSpMaintainVtwoCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is MakeMaintainPlanTopRequest Setter
// 维修方案
func (r *CainiaoIotTicketSpMaintainVtwoCreateAPIRequest) SetMakeMaintainPlanTopRequest(_makeMaintainPlanTopRequest *MakeMaintainPlanV2TopRequest) error {
	r._makeMaintainPlanTopRequest = _makeMaintainPlanTopRequest
	r.Set("make_maintain_plan_top_request", _makeMaintainPlanTopRequest)
	return nil
}

// Get MakeMaintainPlanTopRequest Getter
func (r CainiaoIotTicketSpMaintainVtwoCreateAPIRequest) GetMakeMaintainPlanTopRequest() *MakeMaintainPlanV2TopRequest {
	return r._makeMaintainPlanTopRequest
}
