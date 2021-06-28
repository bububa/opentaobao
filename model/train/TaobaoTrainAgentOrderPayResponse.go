package train

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
代购订单代付接口 APIResponse
taobao.train.agent.order.pay

代购订单代付接口
*/
type TaobaoTrainAgentOrderPayAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"train_agent_order_pay_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 成功返回
    
    ExtendParams   string `json:"extend_params,omitempty" xml:"