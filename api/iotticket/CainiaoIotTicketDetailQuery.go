package iotticket

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iotticket"
)

// CainiaoIotTicketDetailQuery IoT售后工单详情查询
// cainiao.iot.ticket.detail.query
//
// Iot售后工单详情信息查询
func CainiaoIotTicketDetailQuery(ctx context.Context, clt *core.SDKClient, req *iotticket.CainiaoIotTicketDetailQueryAPIRequest, resp *iotticket.CainiaoIotTicketDetailQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
