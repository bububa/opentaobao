package iotticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoIotTicketSpMailSignUploadAPIRequest
IoT售后服务商签收客户邮寄设备附件上传 API请求
cainiao.iot.ticket.sp.mail.sign.upload

IoT售后服务商签收客户邮寄设备附件上传 */
type CainiaoIotTicketSpMailSignUploadAPIRequest struct {
	model.Params
	// 请求参数
	_param *UploadSignVoucherRequest
}

// New
