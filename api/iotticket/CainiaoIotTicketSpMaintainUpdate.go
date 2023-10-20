package iotticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iotticket"
)

// Cainiaoiotticketspmaintainupdate IoT售后服务商维修方案更新
// cainiao.iot.ticket.sp.maintain.update
//
// IoT售后服务商维修方案更新
func Cainiaoiotticketspmaintainupdate(clt *core.SDKClient, req *iotticket.CainiaoiotticketspmaintainupdateAPIRequest, session string) (*iotticket.CainiaoiotticketspmaintainupdateAPIResponse, error) {
	var resp iotticket.CainiaoiotticketspmaintainupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
