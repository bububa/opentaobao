package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
产品规格审核信息获取接口 APIResponse
tmall.product.specs.ticket.get

批量根据specId查询产品规格审核信息包括产品规格状态，申请人，拒绝原因等
*/
type TmallProductSpecsTicketGetAPIResponse struct {
    model.CommonResponse
    Response *TmallProductSpecsTicketGetResponse `json:"tmall_product_specs_ticket_get_response,omitempty"`
}

type TmallProductSpecsTicketGetResponse struct {

    // 产品规格审核单信息
    Tickets   []Ticket `json:"tickets,omitempty"`

}
