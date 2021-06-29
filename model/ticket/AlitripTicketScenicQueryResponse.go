package ticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【门票API2.0】卖家已发布门票商品列表查询接口（根据景点维度查询） APIResponse
alitrip.ticket.scenic.query

查询卖家已发布过的门票商品列表，根据景点维度聚合查询。如果卖家在该景点下未发布过任何商品，则查询不到数据！
*/
type AlitripTicketScenicQueryAPIResponse struct {
    model.CommonResponse
    AlitripTicketScenicQueryResponse
}

type AlitripTicketScenicQueryResponse struct {
    XMLName xml.Name `xml:"alitrip_ticket_scenic_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    FirstResult   *ScenicAndProductResult `json:"first_result,omitempty" xml:"first_result,omitempty"`

    
}
