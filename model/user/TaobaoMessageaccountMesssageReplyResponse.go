package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
消息号下行回复接口 APIResponse
taobao.messageaccount.messsage.reply

外部 isv 调用该进口来进行消息号消息的回复
*/
type TaobaoMessageaccountMesssageReplyAPIResponse struct {
    model.CommonResponse
    Response *TaobaoMessageaccountMesssageReplyResponse `json:"taobao_messageaccount_messsage_reply_response,omitempty"`
}

type TaobaoMessageaccountMesssageReplyResponse struct {

    // result
    Result   *TaobaoMessageaccountMesssageReplyResult `json:"result,omitempty"`

}
