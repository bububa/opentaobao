package iotticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoIotTicketSpAcceptAPIRequest
IoT售后服务商确认接单 API请求
cainiao.iot.ticket.sp.accept

IoT售后服务商确认接单 */
type CainiaoIotTicketSpAcceptAPIRequest struct {
	model.Params
	// 请求参数
	_param *AcceptTicketTopRequest
}

// New
