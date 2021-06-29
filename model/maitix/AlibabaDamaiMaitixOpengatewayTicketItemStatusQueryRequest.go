package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销状态查询接口queryTicketItemStatusByTicketItemId APIRequest
alibaba.damai.maitix.opengateway.ticketItem.status.query

queryTicketItemStatusByTicketItemId
*/
type AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryRequest struct {
    model.Params

    // 入参
    disTicketItemStatusQueryParam   *DisTicketItemStatusQueryDto 

}

func NewAlibabaDamaiMaitixOpengatewayTicketItemStatusQueryRequest() *AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryRequest{
    return &AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.opengateway.ticketItem.status.query"
}

func (r AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryRequest) SetDisTicketItemStatusQueryParam(disTicketItemStatusQueryParam *DisTicketItemStatusQueryDto) error {
    r.disTicketItemStatusQueryParam = disTicketItemStatusQueryParam
    r.Set("dis_ticket_item_status_query_param", disTicketItemStatusQueryParam)
    return nil
}

func (r AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryRequest) GetDisTicketItemStatusQueryParam() *DisTicketItemStatusQueryDto {
    return r.disTicketItemStatusQueryParam
}

