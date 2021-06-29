package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销状态查询接口queryTicketItemStatusByTicketItemId API请求
alibaba.damai.maitix.opengateway.ticketItem.status.query

queryTicketItemStatusByTicketItemId
*/
type AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryRequest struct {
    model.Params
    // 入参
    _disTicketItemStatusQueryParam   *DisTicketItemStatusQueryDto
}

// 初始化AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryRequest对象
func NewAlibabaDamaiMaitixOpengatewayTicketItemStatusQueryRequest() *AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryRequest{
    return &AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.opengateway.ticketItem.status.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DisTicketItemStatusQueryParam Setter
// 入参
func (r *AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryRequest) SetDisTicketItemStatusQueryParam(_disTicketItemStatusQueryParam *DisTicketItemStatusQueryDto) error {
    r._disTicketItemStatusQueryParam = _disTicketItemStatusQueryParam
    r.Set("dis_ticket_item_status_query_param", _disTicketItemStatusQueryParam)
    return nil
}

// DisTicketItemStatusQueryParam Getter
func (r AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryRequest) GetDisTicketItemStatusQueryParam() *DisTicketItemStatusQueryDto {
    return r._disTicketItemStatusQueryParam
}