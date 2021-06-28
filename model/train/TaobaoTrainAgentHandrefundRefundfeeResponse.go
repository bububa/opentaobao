package train

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商手动退款接口 APIResponse
taobao.train.agent.handrefund.refundfee

火车票代理商手动退款接口
*/
type TaobaoTrainAgentHandrefundRefundfeeAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"train_agent_handrefund_refundfee_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功标记
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"