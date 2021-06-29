package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
给注册用户发送消息 APIResponse
alibaba.iwork.mc.msg.senddefault

给神鲸注册用户发送对应操作结果的消息
*/
type AlibabaIworkMcMsgSenddefaultAPIResponse struct {
    model.CommonResponse
    AlibabaIworkMcMsgSenddefaultResponse
}

type AlibabaIworkMcMsgSenddefaultResponse struct {
    XMLName xml.Name `xml:"alibaba_iwork_mc_msg_senddefault_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
