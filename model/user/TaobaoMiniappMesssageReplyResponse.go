package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
轻店铺下行回复接口 APIResponse
taobao.miniapp.messsage.reply

外部 isv 调用该进口来进行轻店铺消息的回复
*/
type TaobaoMiniappMesssageReplyAPIResponse struct {
    model.CommonResponse
    Response *TaobaoMiniappMesssageReplyResponse `json:"taobao_miniapp_messsage_reply_response,omitempty"`
}

type TaobaoMiniappMesssageReplyResponse struct {

    // result
    Result   *TaobaoMiniappMesssageReplyResult `json:"result,omitempty"`

}
