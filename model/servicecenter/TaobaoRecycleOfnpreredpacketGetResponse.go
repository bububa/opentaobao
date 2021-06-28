package servicecenter

import (
    "github.com/bububa/opentaobao/model"
)

/* 
服务商查询前置补贴红包的最新数据 APIResponse
taobao.recycle.ofnpreredpacket.get

服务商查询前置补贴红包的最新数据
*/
type TaobaoRecycleOfnpreredpacketGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoRecycleOfnpreredpacketGetResponse `json:"recycle_ofnpreredpacket_get_response,omitempty"` 
    TaobaoRecycleOfnpreredpacketGetResponse
}

/* model for simplify = false
type TaobaoRecycleOfnpreredpacketGetResponse struct {

    // 前置补贴红包
    
    Data  *struct {
        OfnPreRedPacketDTO  *OfnPreRedPacketDTO `json:"ofn_pre_red_packet_dto,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

type TaobaoRecycleOfnpreredpacketGetResponse struct {

    // 前置补贴红包
    Data   *OfnPreRedPacketDTO `json:"data,omitempty"`

}
