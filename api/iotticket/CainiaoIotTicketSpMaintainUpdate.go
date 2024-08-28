package iotticket

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iotticket"
)

// CainiaoIotTicketSpMaintainUpdate IoT售后服务商维修方案更新
// cainiao.iot.ticket.sp.maintain.update
//
// IoT售后服务商维修方案更新
func CainiaoIotTicketSpMaintainUpdate(ctx context.Context, clt *core.SDKClient, req *iotticket.CainiaoIotTicketSpMaintainUpdateAPIRequest, resp *iotticket.CainiaoIotTicketSpMaintainUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
