package iotticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoIotTicketSpMaintainUpdateAPIRequest IoT售后服务商维修方案更新 API请求
// cainiao.iot.ticket.sp.maintain.update
//
// IoT售后服务商维修方案更新
type CainiaoIotTicketSpMaintainUpdateAPIRequest struct {
	model.Params
	// 请求参数
	_param *UpdateMaintainPlanTopRequest
}

// NewCainiaoIotTicketSpMaintainUpdateRequest 初始化CainiaoIotTicketSpMaintainUpdateAPIRequest对象
func NewCainiaoIotTicketSpMaintainUpdateRequest() *CainiaoIotTicketSpMaintainUpdateAPIRequest {
	return &CainiaoIotTicketSpMaintainUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoIotTicketSpMaintainUpdateAPIRequest) GetApiMethodName() string {
	return "cainiao.iot.ticket.sp.maintain.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoIotTicketSpMaintainUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam is Param Setter
// 请求参数
func (r *CainiaoIotTicketSpMaintainUpdateAPIRequest) SetParam(_param *UpdateMaintainPlanTopRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r CainiaoIotTicketSpMaintainUpdateAPIRequest) GetParam() *UpdateMaintainPlanTopRequest {
	return r._param
}
