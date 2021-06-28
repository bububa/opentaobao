package servicecenter

import (
    "github.com/bububa/opentaobao/model"
)

/* 
回收商同步前置补贴红包的代扣成功事件 APIResponse
taobao.recycle.ofnpreredpacket.tpdeductsuccess

回收商->天猫后端，同步前置补贴红包的代扣成功事件
*/
type TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoRecycleOfnpreredpacketTpdeductsuccessResponse `json:"recycle_ofnpreredpacket_tpdeductsuccess_response,omitempty"` 
    TaobaoRecycleOfnpreredpacketTpdeductsuccessResponse
}

/* model for simplify = false
type TaobaoRecycleOfnpreredpacketTpdeductsuccessResponse struct {

    // 操作
    
    Data  *struct {
        OfnPreRedPacketActionDTO  *OfnPreRedPacketActionDTO `json:"ofn_pre_red_packet_action_dto,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

type TaobaoRecycleOfnpreredpacketTpdeductsuccessResponse struct {

    // 操作
    Data   *OfnPreRedPacketActionDTO `json:"data,omitempty"`

}
