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
    TaobaoTrainAgentOrderPayResponse
}

type TaobaoTrainAgentOrderPayResponse struct {
    XMLName xml.Name `xml:"train_agent_order_pay_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 成功返回
    
    ExtendParams   string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`

    
}
