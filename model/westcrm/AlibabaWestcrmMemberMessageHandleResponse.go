package westcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
处理Usc会员消息接口 APIResponse
alibaba.westcrm.member.message.handle

处理Usc会员消息接口
*/
type AlibabaWestcrmMemberMessageHandleAPIResponse struct {
    model.CommonResponse
    AlibabaWestcrmMemberMessageHandleResponse
}

type AlibabaWestcrmMemberMessageHandleResponse struct {
    XMLName xml.Name `xml:"alibaba_westcrm_member_message_handle_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回对象封装
    
    Result   *DataResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
