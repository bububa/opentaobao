package iotticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoIotTicketSpAcceptAPIRequest IoT售后服务商确认接单 API请求
// cainiao.iot.ticket.sp.accept
//
// IoT售后服务商确认接单
type CainiaoIotTicketSpAcceptAPIRequest struct {
	model.Params
	// 请求参数
	_param *AcceptTicketTopRequest
}

// NewCainiaoIotTicketSpAcceptRequest 初始化CainiaoIotTicketSpAcceptAPIRequest对象
func NewCainiaoIotTicketSpAcceptRequest() *CainiaoIotTicketSpAcceptAPIRequest {
	return &CainiaoIotTicketSpAcceptAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoIotTicketSpAcceptAPIRequest) GetApiMethodName() string {
	return "cainiao.iot.ticket.sp.accept"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoIotTicketSpAcceptAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param Setter
// 请求参数
func (r *CainiaoIotTicketSpAcceptAPIRequest) SetParam(_param *AcceptTicketTopRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r CainiaoIotTicketSpAcceptAPIRequest) GetParam() *AcceptTicketTopRequest {
	return r._param
}
