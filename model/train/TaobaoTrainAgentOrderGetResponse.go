package train

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商获取订单信息回调API APIResponse
taobao.train.agent.order.get

代理商获取订单信息回调API
*/
type TaobaoTrainAgentOrderGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"train_agent_order_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // ttp_order_id
    
    TtpOrderId   int64 `json:"ttp_order_id,omitempty" xml:"