package iotticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iotticket"
)

// Cainiaoiotticketdetailquery IoT售后工单详情查询
// cainiao.iot.ticket.detail.query
//
// Iot售后工单详情信息查询
func Cainiaoiotticketdetailquery(clt *core.SDKClient, req *iotticket.CainiaoiotticketdetailqueryAPIRequest, session string) (*iotticket.CainiaoiotticketdetailqueryAPIResponse, error) {
	var resp iotticket.CainiaoiotticketdetailqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
