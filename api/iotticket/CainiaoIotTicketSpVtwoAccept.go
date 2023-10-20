package iotticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iotticket"
)

// CainiaoIotTicketSpVtwoAccept IoT售后服务商确认接单
// cainiao.iot.ticket.sp.vtwo.accept
//
// IoT售后服务商确认接单
func CainiaoIotTicketSpVtwoAccept(clt *core.SDKClient, req *iotticket.CainiaoIotTicketSpVtwoAcceptAPIRequest, resp *iotticket.CainiaoIotTicketSpVtwoAcceptAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
