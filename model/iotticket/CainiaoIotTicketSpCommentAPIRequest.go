package iotticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoIotTicketSpCommentAPIRequest
IoT售后服务商工单备注 API请求
cainiao.iot.ticket.sp.comment

IoT售后服务商工单备注 */
type CainiaoIotTicketSpCommentAPIRequest struct {
	model.Params
	// 请求参数
	_param *CommentTicketTopRequest
}

// New
