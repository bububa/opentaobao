package iotticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoIotTicketSpMailVoucherUploadAPIRequest
服务商寄出维修件上传凭证信息 API请求
cainiao.iot.ticket.sp.mail.voucher.upload

IoT售后服务商寄出维修件上传凭证信息 */
type CainiaoIotTicketSpMailVoucherUploadAPIRequest struct {
	model.Params
	// 请求参数
	_param *CommentTicketTopRequest
}

// New
