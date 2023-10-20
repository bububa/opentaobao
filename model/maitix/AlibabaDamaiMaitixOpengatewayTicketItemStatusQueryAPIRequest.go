package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIRequest 分销状态查询接口queryTicketItemStatusByTicketItemId API请求
// alibaba.damai.maitix.opengateway.ticketItem.status.query
//
// queryTicketItemStatusByTicketItemId
type AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIRequest struct {
	model.Params
	// 入参
	_disTicketItemStatusQueryParam *DisTicketItemStatusQueryDto
}

// NewAlibabaDamaiMaitixOpengatewayTicketItemStatusQueryRequest 初始化AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIRequest对象
func NewAlibabaDamaiMaitixOpengatewayTicketItemStatusQueryRequest() *AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIRequest {
	return &AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.opengateway.ticketItem.status.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDisTicketItemStatusQueryParam is DisTicketItemStatusQueryParam Setter
// 入参
func (r *AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIRequest) SetDisTicketItemStatusQueryParam(_disTicketItemStatusQueryParam *DisTicketItemStatusQueryDto) error {
	r._disTicketItemStatusQueryParam = _disTicketItemStatusQueryParam
	r.Set("dis_ticket_item_status_query_param", _disTicketItemStatusQueryParam)
	return nil
}

// GetDisTicketItemStatusQueryParam DisTicketItemStatusQueryParam Getter
func (r AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIRequest) GetDisTicketItemStatusQueryParam() *DisTicketItemStatusQueryDto {
	return r._disTicketItemStatusQueryParam
}
