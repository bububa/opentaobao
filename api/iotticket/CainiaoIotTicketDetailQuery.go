package iotticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iotticket"
)

// CainiaoIotTicketDetailQuery IoT售后工单详情查询
// cainiao.iot.ticket.detail.query
//
// Iot售后工单详情信息查询
func CainiaoIotTicketDetailQuery(clt *core.SDKClient, req *iotticket.CainiaoIotTicketDetailQueryAPIRequest, session string) (*iotticket.CainiaoIotTicketDetailQueryAPIResponse, error) {
	var resp iotticket.CainiaoIotTicketDetailQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
