package train

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
火车票代理商接口——确认出票是否成功 API返回值 
taobao.train.agent.bookticket.confirm

火车票代理商接口——确认出票是否成功
*/
type TaobaoTrainAgentBookticketConfirmAPIResponse struct {
    model.CommonResponse
    TaobaoTrainAgentBookticketConfirmAPIResponseModel
}

// 火车票代理商接口——确认出票是否成功 成功返回结果
type TaobaoTrainAgentBookticketConfirmAPIResponseModel struct {
    XMLName xml.Name `xml:"train_agent_bookticket_confirm_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
