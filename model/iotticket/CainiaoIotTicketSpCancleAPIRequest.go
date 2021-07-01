package iotticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoIotTicketSpCancleAPIRequest
Iot售后服务商取消工单 API请求
cainiao.iot.ticket.sp.cancle

IoT售后服务商取消接单 */
type CainiaoIotTicketSpCancleAPIRequest struct {
	model.Params
	// 请求参数
	_param *AcceptTicketTopRequest
}

// NewCainiaoIotTicketSpCancleRequest 初始化CainiaoIotTicketSpCancleAPIRequest对象
func NewCainiaoIotTicketSpCancleRequest() *CainiaoIotTicketSpCancleAPIRequest {
	return &CainiaoIotTicketSpCancleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoIotTicketSpCancleAPIRequest) GetApiMethodName() string {
	return "cainiao.iot.ticket.sp.cancle"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoIotTicketSpCancleAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param Setter
// 请求参数
func (r *CainiaoIotTicketSpCancleAPIRequest) SetParam(_param *AcceptTicketTopRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r CainiaoIotTicketSpCancleAPIRequest) GetParam() *AcceptTicketTopRequest {
	return r._param
}
