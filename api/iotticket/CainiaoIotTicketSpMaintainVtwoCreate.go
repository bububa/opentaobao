package iotticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iotticket"
)

// CainiaoIotTicketSpMaintainVtwoCreate 服务商制定维修费方案
// cainiao.iot.ticket.sp.maintain.vtwo.create
//
// 服务商制定维修费方案
func CainiaoIotTicketSpMaintainVtwoCreate(clt *core.SDKClient, req *iotticket.CainiaoIotTicketSpMaintainVtwoCreateAPIRequest, resp *iotticket.CainiaoIotTicketSpMaintainVtwoCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
