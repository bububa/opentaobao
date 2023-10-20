package iotticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iotticket"
)

// CainiaoIotTicketSpMaintainUpdate IoT售后服务商维修方案更新
// cainiao.iot.ticket.sp.maintain.update
//
// IoT售后服务商维修方案更新
func CainiaoIotTicketSpMaintainUpdate(clt *core.SDKClient, req *iotticket.CainiaoIotTicketSpMaintainUpdateAPIRequest, resp *iotticket.CainiaoIotTicketSpMaintainUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
