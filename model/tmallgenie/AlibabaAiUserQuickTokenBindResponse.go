package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
人工智能实验室精灵用户绑定第三方Token接口 APIResponse
alibaba.ai.user.quick.token.bind

人工智能实验室精灵用户绑定第三方Token接口
*/
type AlibabaAiUserQuickTokenBindAPIResponse struct {
    model.CommonResponse
    AlibabaAiUserQuickTokenBindResponse
}

type AlibabaAiUserQuickTokenBindResponse struct {
    XMLName xml.Name `xml:"alibaba_ai_user_quick_token_bind_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // message
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // statusCode
    
    StatusCode   int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`

    
}
