package iotticket

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iotticket"
)

// CainiaoIotTicketSpVtwoAccept IoT售后服务商确认接单
// cainiao.iot.ticket.sp.vtwo.accept
//
// IoT售后服务商确认接单
func CainiaoIotTicketSpVtwoAccept(ctx context.Context, clt *core.SDKClient, req *iotticket.CainiaoIotTicketSpVtwoAcceptAPIRequest, resp *iotticket.CainiaoIotTicketSpVtwoAcceptAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
