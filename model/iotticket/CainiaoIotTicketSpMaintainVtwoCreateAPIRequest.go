package iotticket

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoIotTicketSpMaintainVtwoCreateAPIRequest) Reset() {
	r._makeMaintainPlanTopRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoIotTicketSpMaintainVtwoCreateAPIRequest) GetApiMethodName() string {
	return "cainiao.iot.ticket.sp.maintain.vtwo.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoIotTicketSpMaintainVtwoCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoIotTicketSpMaintainVtwoCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMakeMaintainPlanTopRequest is MakeMaintainPlanTopRequest Setter
// 维修方案
func (r *CainiaoIotTicketSpMaintainVtwoCreateAPIRequest) SetMakeMaintainPlanTopRequest(_makeMaintainPlanTopRequest *MakeMaintainPlanV2TopRequest) error {
	r._makeMaintainPlanTopRequest = _makeMaintainPlanTopRequest
	r.Set("make_maintain_plan_top_request", _makeMaintainPlanTopRequest)
	return nil
}

// GetMakeMaintainPlanTopRequest MakeMaintainPlanTopRequest Getter
func (r CainiaoIotTicketSpMaintainVtwoCreateAPIRequest) GetMakeMaintainPlanTopRequest() *MakeMaintainPlanV2TopRequest {
	return r._makeMaintainPlanTopRequest
}

var poolCainiaoIotTicketSpMaintainVtwoCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoIotTicketSpMaintainVtwoCreateRequest()
	},
}

// GetCainiaoIotTicketSpMaintainVtwoCreateRequest 从 sync.Pool 获取 CainiaoIotTicketSpMaintainVtwoCreateAPIRequest
func GetCainiaoIotTicketSpMaintainVtwoCreateAPIRequest() *CainiaoIotTicketSpMaintainVtwoCreateAPIRequest {
	return poolCainiaoIotTicketSpMaintainVtwoCreateAPIRequest.Get().(*CainiaoIotTicketSpMaintainVtwoCreateAPIRequest)
}

// ReleaseCainiaoIotTicketSpMaintainVtwoCreateAPIRequest 将 CainiaoIotTicketSpMaintainVtwoCreateAPIRequest 放入 sync.Pool
func ReleaseCainiaoIotTicketSpMaintainVtwoCreateAPIRequest(v *CainiaoIotTicketSpMaintainVtwoCreateAPIRequest) {
	v.Reset()
	poolCainiaoIotTicketSpMaintainVtwoCreateAPIRequest.Put(v)
}
