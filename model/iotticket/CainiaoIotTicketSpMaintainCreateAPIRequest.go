package iotticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoIotTicketSpMaintainCreateAPIRequest
IoT售后服务商制定维修方案 API请求
cainiao.iot.ticket.sp.maintain.create

IoT售后服务商制定维修方案 */
type CainiaoIotTicketSpMaintainCreateAPIRequest struct {
	model.Params
	// 请求参数
	_param *AssignMaintenancePersonnelTopRequest
}

// NewCainiaoIotTicketSpMaintainCreateRequest 初始化CainiaoIotTicketSpMaintainCreateAPIRequest对象
func NewCainiaoIotTicketSpMaintainCreateRequest() *CainiaoIotTicketSpMaintainCreateAPIRequest {
	return &CainiaoIotTicketSpMaintainCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoIotTicketSpMaintainCreateAPIRequest) GetApiMethodName() string {
	return "cainiao.iot.ticket.sp.maintain.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoIotTicketSpMaintainCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param Setter
// 请求参数
func (r *CainiaoIotTicketSpMaintainCreateAPIRequest) SetParam(_param *AssignMaintenancePersonnelTopRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r CainiaoIotTicketSpMaintainCreateAPIRequest) GetParam() *AssignMaintenancePersonnelTopRequest {
	return r._param
}
