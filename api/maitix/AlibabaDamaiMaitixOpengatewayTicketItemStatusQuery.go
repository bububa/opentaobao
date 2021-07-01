package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

/* AlibabaDamaiMaitixOpengatewayTicketItemStatusQuery
分销状态查询接口queryTicketItemStatusByTicketItemId
alibaba.damai.maitix.opengateway.ticketItem.status.query

queryTicketItemStatusByTicketItemId */
func AlibabaDamaiMaitixOpengatewayTicketItemStatusQuery(clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIRequest, session string) (*maitix.AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIResponse, error) {
	var resp maitix.AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
