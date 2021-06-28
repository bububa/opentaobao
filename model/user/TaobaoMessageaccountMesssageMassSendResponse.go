package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
消息号开放-消息群发 APIResponse
taobao.messageaccount.messsage.mass.send

外部 isv 调用该进口来进行消息号消息的群发
*/
type TaobaoMessageaccountMesssageMassSendAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoMessageaccountMesssageMassSendResponse `json:"messageaccount_messsage_mass_send_response,omitempty"` 
    TaobaoMessageaccountMesssageMassSendResponse
}

/* model for simplify = false
type TaobaoMessageaccountMesssageMassSendResponse struct {

    // result
    
    Result  *struct {
        TaobaoMessageaccountMesssageMassSendResult  *TaobaoMessageaccountMesssageMassSendResult `json:"taobao_messageaccount_messsage_mass_send_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoMessageaccountMesssageMassSendResponse struct {

    // result
    Result   *TaobaoMessageaccountMesssageMassSendResult `json:"result,omitempty"`

}
