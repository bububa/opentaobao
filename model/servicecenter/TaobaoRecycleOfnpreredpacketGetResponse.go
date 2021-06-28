package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商查询前置补贴红包的最新数据 APIResponse
taobao.recycle.ofnpreredpacket.get

服务商查询前置补贴红包的最新数据
*/
type TaobaoRecycleOfnpreredpacketGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"recycle_ofnpreredpacket_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 前置补贴红包
    
    Data   *OfnPreRedPacketDTO `json:"data,omitempty" xml:"