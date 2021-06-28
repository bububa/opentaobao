package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
回收商同步前置补贴红包的代扣成功事件 APIResponse
taobao.recycle.ofnpreredpacket.tpdeductsuccess

回收商->天猫后端，同步前置补贴红包的代扣成功事件
*/
type TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"recycle_ofnpreredpacket_tpdeductsuccess_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 操作
    
    Data   *OfnPreRedPacketActionDTO `json:"data,omitempty" xml:"