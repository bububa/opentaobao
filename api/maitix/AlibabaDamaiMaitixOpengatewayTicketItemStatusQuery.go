package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// AlibabadamaimaitixopengatewayticketItemstatusquery 分销状态查询接口queryTicketItemStatusByTicketItemId
// alibaba.damai.maitix.opengateway.ticketItem.status.query
//
// queryTicketItemStatusByTicketItemId
func AlibabadamaimaitixopengatewayticketItemstatusquery(clt *core.SDKClient, req *maitix.AlibabadamaimaitixopengatewayticketItemstatusqueryAPIRequest, session string) (*maitix.AlibabadamaimaitixopengatewayticketItemstatusqueryAPIResponse, error) {
	var resp maitix.AlibabadamaimaitixopengatewayticketItemstatusqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
