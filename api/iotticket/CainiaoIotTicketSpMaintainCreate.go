package iotticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iotticket"
)

// CainiaoIotTicketSpMaintainCreate IoT售后服务商制定维修方案
// cainiao.iot.ticket.sp.maintain.create
//
// IoT售后服务商制定维修方案
func CainiaoIotTicketSpMaintainCreate(clt *core.SDKClient, req *iotticket.CainiaoIotTicketSpMaintainCreateAPIRequest, session string) (*iotticket.CainiaoIotTicketSpMaintainCreateAPIResponse, error) {
	var resp iotticket.CainiaoIotTicketSpMaintainCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
