package ticket

import (
    "github.com/bububa/opentaobao/model"
)

/* 
【门票API2.0】卖家已发布门票商品列表查询接口（根据景点维度查询） APIResponse
alitrip.ticket.scenic.query

查询卖家已发布过的门票商品列表，根据景点维度聚合查询。如果卖家在该景点下未发布过任何商品，则查询不到数据！
*/
type AlitripTicketScenicQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlitripTicketScenicQueryResponse `json:"alitrip_ticket_scenic_query_response,omitempty"` 
    AlitripTicketScenicQueryResponse
}

/* model for simplify = false
type AlitripTicketScenicQueryResponse struct {

    // 返回结果
    
    FirstResult  *struct {
        ScenicAndProductResult  *ScenicAndProductResult `json:"scenic_and_product_result,omitempty"`
    } `json:"first_result,omitempty"`
    

}
*/

type AlitripTicketScenicQueryResponse struct {

    // 返回结果
    FirstResult   *ScenicAndProductResult `json:"first_result,omitempty"`

}
