package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
产品规格审核信息获取接口 API返回值 
tmall.product.specs.ticket.get

批量根据specId查询产品规格审核信息包括产品规格状态，申请人，拒绝原因等
*/
type TmallProductSpecsTicketGetAPIResponse struct {
    model.CommonResponse
    TmallProductSpecsTicketGetResponse
}

// 产品规格审核信息获取接口 成功返回结果
type TmallProductSpecsTicketGetResponse struct {
    XMLName xml.Name `xml:"tmall_product_specs_ticket_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 产品规格审核单信息
    Tickets   []Ticket `json:"tickets,omitempty" xml:"tickets>ticket,omitempty"`
}
