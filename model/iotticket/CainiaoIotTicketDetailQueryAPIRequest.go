package iotticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoIotTicketDetailQueryAPIRequest
IoT售后工单详情查询 API请求
cainiao.iot.ticket.detail.query

Iot售后工单详情信息查询 */
type CainiaoIotTicketDetailQueryAPIRequest struct {
	model.Params
	// 服务商唯一编码
	_spCode string
	// 工单Id
	_ticketId int64
}

// New
