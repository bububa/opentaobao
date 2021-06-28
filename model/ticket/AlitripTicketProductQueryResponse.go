package ticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【门票API2.0】门票商品查询接口 APIResponse
alitrip.ticket.product.query

门票商品查询接口：返回商家上传的门票商品信息
*/
type AlitripTicketProductQueryAPIResponse struct {
    model.CommonResponse
    AlitripTicketProductQueryResponse
}

type AlitripTicketProductQueryResponse struct {
    XMLName xml.Name `xml:"alitrip_ticket_product_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 门票商品详情
    
    FirstResult   *TopTicketItemFullinfoResult `json:"first_result,omitempty" xml:"first_result,omitempty"`

    
}
