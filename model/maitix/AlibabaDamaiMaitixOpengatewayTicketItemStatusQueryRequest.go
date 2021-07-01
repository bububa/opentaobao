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
type AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIRequest struct {
    model.Params
    // 入参
    _disTicketItemStatusQueryParam   *DisTicketItemStatusQueryDTO
}

// 初始化AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIRequest对象
func NewAlibabaDamaiMaitixOpengatewayTicketItemStatusQueryRequest() *AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIRequest{
    return &AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.opengateway.ticketItem.status.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DisTicketItemStatusQueryParam Setter
// 入参
func (r *AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIRequest) SetDisTicketItemStatusQueryParam(_disTicketItemStatusQueryParam *DisTicketItemStatusQueryDTO) error {
    r._disTicketItemStatusQueryParam = _disTicketItemStatusQueryParam
    r.Set("dis_ticket_item_status_query_param", _disTicketItemStatusQueryParam)
    return nil
}

// DisTicketItemStatusQueryParam Getter
func (r AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIRequest) GetDisTicketItemStatusQueryParam() *DisTicketItemStatusQueryDTO {
    return r._disTicketItemStatusQueryParam
}
