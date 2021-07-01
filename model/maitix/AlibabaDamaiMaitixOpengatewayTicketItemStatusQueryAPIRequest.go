package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIRequest
分销状态查询接口queryTicketItemStatusByTicketItemId API请求
alibaba.damai.maitix.opengateway.ticketItem.status.query

queryTicketItemStatusByTicketItemId */
type AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIRequest struct {
	model.Params
	// 入参
	_disTicketItemStatusQueryParam *DisTicketItemStatusQueryDto
}

// New
