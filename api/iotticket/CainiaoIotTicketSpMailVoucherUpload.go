package iotticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iotticket"
)

// CainiaoIotTicketSpMailVoucherUpload 服务商寄出维修件上传凭证信息
// cainiao.iot.ticket.sp.mail.voucher.upload
//
// IoT售后服务商寄出维修件上传凭证信息
func CainiaoIotTicketSpMailVoucherUpload(clt *core.SDKClient, req *iotticket.CainiaoIotTicketSpMailVoucherUploadAPIRequest, session string) (*iotticket.CainiaoIotTicketSpMailVoucherUploadAPIResponse, error) {
	var resp iotticket.CainiaoIotTicketSpMailVoucherUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
