package iotticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iotticket"
)

// Cainiaoiotticketspmailsignupload IoT售后服务商签收客户邮寄设备附件上传
// cainiao.iot.ticket.sp.mail.sign.upload
//
// IoT售后服务商签收客户邮寄设备附件上传
func Cainiaoiotticketspmailsignupload(clt *core.SDKClient, req *iotticket.CainiaoiotticketspmailsignuploadAPIRequest, session string) (*iotticket.CainiaoiotticketspmailsignuploadAPIResponse, error) {
	var resp iotticket.CainiaoiotticketspmailsignuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
