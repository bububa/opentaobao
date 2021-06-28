package train

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
火车票代理商接口——订单关闭实际出票成功审计接口 APIResponse
taobao.train.agent.direct.compensate

代购直连订单平台关单但是代理商出票成功补偿接口
*/
type TaobaoTrainAgentDirectCompensateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"train_agent_direct_compensate_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"