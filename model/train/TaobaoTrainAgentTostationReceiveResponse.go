package train

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
线下票送票至车站代理商确认用户已取票服务 APIResponse
taobao.train.agent.tostation.receive

送票至车站的订单，代理商确认用户已取票
*/
type TaobaoTrainAgentTostationReceiveAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"train_agent_tostation_receive_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误码
    
    ErrorMsgCode   string `json:"error_msg_code,omitempty" xml:"