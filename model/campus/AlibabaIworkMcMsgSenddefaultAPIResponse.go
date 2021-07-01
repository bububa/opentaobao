package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
给注册用户发送消息 API返回值 
alibaba.iwork.mc.msg.senddefault

给神鲸注册用户发送对应操作结果的消息
*/
type AlibabaIworkMcMsgSenddefaultAPIResponse struct {
    model.CommonResponse
    AlibabaIworkMcMsgSenddefaultAPIResponseModel
}

// 给注册用户发送消息 成功返回结果
type AlibabaIworkMcMsgSenddefaultAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_iwork_mc_msg_senddefault_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
