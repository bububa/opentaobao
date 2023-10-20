package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimaitixopengatewayticketItemstatusqueryAPIRequest 分销状态查询接口queryTicketItemStatusByTicketItemId API请求
// alibaba.damai.maitix.opengateway.ticketItem.status.query
//
// queryTicketItemStatusByTicketItemId
type AlibabadamaimaitixopengatewayticketItemstatusqueryAPIRequest struct {
	model.Params
	// 入参
	_disTicketItemStatusQueryParam *DisTicketItemStatusQueryDto
}

// NewAlibabadamaimaitixopengatewayticketItemstatusqueryRequest 初始化AlibabadamaimaitixopengatewayticketItemstatusqueryAPIRequest对象
func NewAlibabadamaimaitixopengatewayticketItemstatusqueryRequest() *AlibabadamaimaitixopengatewayticketItemstatusqueryAPIRequest {
	return &AlibabadamaimaitixopengatewayticketItemstatusqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimaitixopengatewayticketItemstatusqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.opengateway.ticketItem.status.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimaitixopengatewayticketItemstatusqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimaitixopengatewayticketItemstatusqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDisTicketItemStatusQueryParam is DisTicketItemStatusQueryParam Setter
// 入参
func (r *AlibabadamaimaitixopengatewayticketItemstatusqueryAPIRequest) SetDisTicketItemStatusQueryParam(_disTicketItemStatusQueryParam *DisTicketItemStatusQueryDto) error {
	r._disTicketItemStatusQueryParam = _disTicketItemStatusQueryParam
	r.Set("dis_ticket_item_status_query_param", _disTicketItemStatusQueryParam)
	return nil
}

// GetDisTicketItemStatusQueryParam DisTicketItemStatusQueryParam Getter
func (r AlibabadamaimaitixopengatewayticketItemstatusqueryAPIRequest) GetDisTicketItemStatusQueryParam() *DisTicketItemStatusQueryDto {
	return r._disTicketItemStatusQueryParam
}
