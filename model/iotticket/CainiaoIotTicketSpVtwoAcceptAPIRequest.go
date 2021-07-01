package iotticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoIotTicketSpVtwoAcceptAPIRequest
IoT售后服务商确认接单 API请求
cainiao.iot.ticket.sp.vtwo.accept

IoT售后服务商确认接单 */
type CainiaoIotTicketSpVtwoAcceptAPIRequest struct {
	model.Params
	// 受理接口请求参数
	_acceptTicketTopRequest *AcceptTicketV2TopRequest
}

// New
