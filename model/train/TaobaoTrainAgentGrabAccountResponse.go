package train

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
代购抢代理商回传12306账号 API返回值 
taobao.train.agent.grab.account

火车票业务代购抢功能，代理商回传12306账号，用于自营抢票链路出票
*/
type TaobaoTrainAgentGrabAccountAPIResponse struct {
    model.CommonResponse
    TaobaoTrainAgentGrabAccountResponse
}

// 代购抢代理商回传12306账号 成功返回结果
type TaobaoTrainAgentGrabAccountResponse struct {
    XMLName xml.Name `xml:"train_agent_grab_account_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // resultMsg
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
    // resultCode
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // success
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
