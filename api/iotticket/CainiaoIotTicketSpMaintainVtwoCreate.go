package iotticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iotticket"
)

// Cainiaoiotticketspmaintainvtwocreate 服务商制定维修费方案
// cainiao.iot.ticket.sp.maintain.vtwo.create
//
// 服务商制定维修费方案
func Cainiaoiotticketspmaintainvtwocreate(clt *core.SDKClient, req *iotticket.CainiaoiotticketspmaintainvtwocreateAPIRequest, session string) (*iotticket.CainiaoiotticketspmaintainvtwocreateAPIResponse, error) {
	var resp iotticket.CainiaoiotticketspmaintainvtwocreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
