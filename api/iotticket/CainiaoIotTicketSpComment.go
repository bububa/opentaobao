package iotticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iotticket"
)

// Cainiaoiotticketspcomment IoT售后服务商工单备注
// cainiao.iot.ticket.sp.comment
//
// IoT售后服务商工单备注
func Cainiaoiotticketspcomment(clt *core.SDKClient, req *iotticket.CainiaoiotticketspcommentAPIRequest, session string) (*iotticket.CainiaoiotticketspcommentAPIResponse, error) {
	var resp iotticket.CainiaoiotticketspcommentAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
