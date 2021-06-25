package ticket

import (
    "github.com/bububa/opentaobao/model"
)

/* 
【门票API2.0】门票景点绑定接口 APIResponse
alitrip.ticket.scenic.bind

门票景点绑定接口，用于建立阿里标准景点id与商家系统景点id的映射关系。该接口同时支持新建和修改映射关系，当用户没有为ali_scenic_id建立过映射关系时，则判断为新建映射关系，否则为修改。可以通过设置update_out_scenic_id来修改ali_scenic_id与out_scenic_id的映射关系。
*/
type AlitripTicketScenicBindAPIResponse struct {
    model.CommonResponse
    Response *AlitripTicketScenicBindResponse `json:"alitrip_ticket_scenic_bind_response,omitempty"`
}

type AlitripTicketScenicBindResponse struct {

    // 景点绑定结果
    FirstResult   *TicketScenicResult `json:"first_result,omitempty"`

}
