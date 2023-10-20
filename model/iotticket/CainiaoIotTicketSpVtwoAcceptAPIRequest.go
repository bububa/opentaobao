package iotticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoiotticketspvtwoacceptAPIRequest IoT售后服务商确认接单 API请求
// cainiao.iot.ticket.sp.vtwo.accept
//
// IoT售后服务商确认接单
type CainiaoiotticketspvtwoacceptAPIRequest struct {
	model.Params
	// 受理接口请求参数
	_acceptTicketTopRequest *AcceptTicketV2topRequest
}

// NewCainiaoiotticketspvtwoacceptRequest 初始化CainiaoiotticketspvtwoacceptAPIRequest对象
func NewCainiaoiotticketspvtwoacceptRequest() *CainiaoiotticketspvtwoacceptAPIRequest {
	return &CainiaoiotticketspvtwoacceptAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoiotticketspvtwoacceptAPIRequest) GetApiMethodName() string {
	return "cainiao.iot.ticket.sp.vtwo.accept"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoiotticketspvtwoacceptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoiotticketspvtwoacceptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAcceptTicketTopRequest is AcceptTicketTopRequest Setter
// 受理接口请求参数
func (r *CainiaoiotticketspvtwoacceptAPIRequest) SetAcceptTicketTopRequest(_acceptTicketTopRequest *AcceptTicketV2topRequest) error {
	r._acceptTicketTopRequest = _acceptTicketTopRequest
	r.Set("accept_ticket_top_request", _acceptTicketTopRequest)
	return nil
}

// GetAcceptTicketTopRequest AcceptTicketTopRequest Getter
func (r CainiaoiotticketspvtwoacceptAPIRequest) GetAcceptTicketTopRequest() *AcceptTicketV2topRequest {
	return r._acceptTicketTopRequest
}
