package train

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
线下票确认送票至车站服务 APIResponse
taobao.train.agent.tostation.confirm

送票至车站的订单，代理商确认配送到站
*/
type TaobaoTrainAgentTostationConfirmAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"train_agent_tostation_confirm_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误码
    
    ErrorMsgCode   string `json:"error_msg_code,omitempty" xml:"