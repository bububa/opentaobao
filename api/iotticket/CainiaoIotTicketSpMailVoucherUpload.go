package iotticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iotticket"
)

// Cainiaoiotticketspmailvoucherupload 服务商寄出维修件上传凭证信息
// cainiao.iot.ticket.sp.mail.voucher.upload
//
// IoT售后服务商寄出维修件上传凭证信息
func Cainiaoiotticketspmailvoucherupload(clt *core.SDKClient, req *iotticket.CainiaoiotticketspmailvoucheruploadAPIRequest, session string) (*iotticket.CainiaoiotticketspmailvoucheruploadAPIResponse, error) {
	var resp iotticket.CainiaoiotticketspmailvoucheruploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
