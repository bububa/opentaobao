package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
下行普通消息 APIResponse
taobao.messageaccount.messsage.normal.send

消息号下行单个普通消息
*/
type TaobaoMessageaccountMesssageNormalSendAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoMessageaccountMesssageNormalSendResponse `json:"messageaccount_messsage_normal_send_response,omitempty"` 
    TaobaoMessageaccountMesssageNormalSendResponse
}

/* model for simplify = false
type TaobaoMessageaccountMesssageNormalSendResponse struct {

    // model
    
    Model   string `json:"model,omitempty"`
    

}
*/

type TaobaoMessageaccountMesssageNormalSendResponse struct {

    // model
    Model   string `json:"model,omitempty"`

}
