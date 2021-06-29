package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
精灵用户绑定第三方账号信息 APIResponse
alibaba.ai.user.quick.bind

人工智能实验室精灵用户绑定第三方账号信息接口，开放给Iot厂商做为厂商上送第三方账号信息的接口
*/
type AlibabaAiUserQuickBindAPIResponse struct {
    model.CommonResponse
    AlibabaAiUserQuickBindResponse
}

type AlibabaAiUserQuickBindResponse struct {
    XMLName xml.Name `xml:"alibaba_ai_user_quick_bind_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 交易结果描述（例如： 交易成功、交易失败）
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 交易结果码（200：成功、其他：失败）
    
    StatusCode   int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`

    
    // 扩展字段，无用
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
    // 交易结果详细描述（例如：用户已经存在，交易失败）
    
    MessageDetail   string `json:"message_detail,omitempty" xml:"message_detail,omitempty"`

    
}
