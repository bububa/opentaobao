package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
下行普通消息 APIResponse
taobao.messageaccount.messsage.normal.send

消息号下行单个普通消息
*/
type TaobaoMessageaccountMesssageNormalSendAPIResponse struct {
    model.CommonResponse
    TaobaoMessageaccountMesssageNormalSendResponse
}

type TaobaoMessageaccountMesssageNormalSendResponse struct {
    XMLName xml.Name `xml:"messageaccount_messsage_normal_send_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // model
    
    Model   string `json:"model,omitempty" xml:"model,omitempty"`

    
}
