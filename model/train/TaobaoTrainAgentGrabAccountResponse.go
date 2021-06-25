package train

import (
    "github.com/bububa/opentaobao/model"
)

/* 
代购抢代理商回传12306账号 APIResponse
taobao.train.agent.grab.account

火车票业务代购抢功能，代理商回传12306账号，用于自营抢票链路出票
*/
type TaobaoTrainAgentGrabAccountAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTrainAgentGrabAccountResponse `json:"taobao_train_agent_grab_account_response,omitempty"`
}

type TaobaoTrainAgentGrabAccountResponse struct {

    // resultMsg
    ResultMsg   string `json:"result_msg,omitempty"`

    // resultCode
    ResultCode   string `json:"result_code,omitempty"`

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

}
