package iotticket

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iotticket"
)

// CainiaoIotTicketSpMailSignUpload IoT售后服务商签收客户邮寄设备附件上传
// cainiao.iot.ticket.sp.mail.sign.upload
//
// IoT售后服务商签收客户邮寄设备附件上传
func CainiaoIotTicketSpMailSignUpload(ctx context.Context, clt *core.SDKClient, req *iotticket.CainiaoIotTicketSpMailSignUploadAPIRequest, resp *iotticket.CainiaoIotTicketSpMailSignUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
