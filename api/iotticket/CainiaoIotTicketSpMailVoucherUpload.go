package iotticket

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iotticket"
)

// CainiaoIotTicketSpMailVoucherUpload 服务商寄出维修件上传凭证信息
// cainiao.iot.ticket.sp.mail.voucher.upload
//
// IoT售后服务商寄出维修件上传凭证信息
func CainiaoIotTicketSpMailVoucherUpload(ctx context.Context, clt *core.SDKClient, req *iotticket.CainiaoIotTicketSpMailVoucherUploadAPIRequest, resp *iotticket.CainiaoIotTicketSpMailVoucherUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
