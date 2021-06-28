package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
消息号开放-消息群发 APIResponse
taobao.messageaccount.messsage.mass.send

外部 isv 调用该进口来进行消息号消息的群发
*/
type TaobaoMessageaccountMesssageMassSendAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"messageaccount_messsage_mass_send_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TaobaoMessageaccountMesssageMassSendResult `json:"result,omitempty" xml:"