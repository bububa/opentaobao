package ticket

import (
    "github.com/bububa/opentaobao/model"
)

/* 
【门票API2.0】门票商品查询接口 APIResponse
alitrip.ticket.product.query

门票商品查询接口：返回商家上传的门票商品信息
*/
type AlitripTicketProductQueryAPIResponse struct {
    model.CommonResponse
    Response *AlitripTicketProductQueryResponse `json:"alitrip_ticket_product_query_response,omitempty"`
}

type AlitripTicketProductQueryResponse struct {

    // 门票商品详情
    FirstResult   *TopTicketItemFullinfoResult `json:"first_result,omitempty"`

}
