package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵消息推送标准接口 APIResponse
alibaba.ailabs.aligenie.skill.message.push

用于AliGenie技能开发者在技能内主动向音响推送消息的标准服务接口，只有订阅过该消息的用户才能收到消息；
*/
type AlibabaAilabsAligenieSkillMessagePushAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsAligenieSkillMessagePushResponse
}

type AlibabaAilabsAligenieSkillMessagePushResponse struct {
    XMLName xml.Name `xml:"alibaba_ailabs_aligenie_skill_message_push_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 消息推送错误码
    
    PushErrorCode   string `json:"push_error_code,omitempty" xml:"push_error_code,omitempty"`

    
    // 消息推送错误信息
    
    PushErrorMsg   string `json:"push_error_msg,omitempty" xml:"push_error_msg,omitempty"`

    
    // 消息推送是否成功
    
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`

    
    // 接口调用是否成功
    
    PushResult   bool `json:"push_result,omitempty" xml:"push_result,omitempty"`

    
}
