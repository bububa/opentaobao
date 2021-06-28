package train

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
代购抢代理商回传12306账号 APIResponse
taobao.train.agent.grab.account

火车票业务代购抢功能，代理商回传12306账号，用于自营抢票链路出票
*/
type TaobaoTrainAgentGrabAccountAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"train_agent_grab_account_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // resultMsg
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"