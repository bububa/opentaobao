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

// New
