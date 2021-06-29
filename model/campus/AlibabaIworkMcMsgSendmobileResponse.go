package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发送消息给手机用户 APIResponse
alibaba.iwork.mc.msg.sendmobile

给手机用户发送对应操作结果的消息
*/
type AlibabaIworkMcMsgSendmobileAPIResponse struct {
    model.CommonResponse
    AlibabaIworkMcMsgSendmobileResponse
}

type AlibabaIworkMcMsgSendmobileResponse struct {
    XMLName xml.Name `xml:"alibaba_iwork_mc_msg_sendmobile_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
