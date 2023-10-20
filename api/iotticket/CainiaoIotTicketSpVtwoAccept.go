package iotticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iotticket"
)

// Cainiaoiotticketspvtwoaccept IoT售后服务商确认接单
// cainiao.iot.ticket.sp.vtwo.accept
//
// IoT售后服务商确认接单
func Cainiaoiotticketspvtwoaccept(clt *core.SDKClient, req *iotticket.CainiaoiotticketspvtwoacceptAPIRequest, session string) (*iotticket.CainiaoiotticketspvtwoacceptAPIResponse, error) {
	var resp iotticket.CainiaoiotticketspvtwoacceptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
