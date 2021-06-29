package alilabs

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取消账号授权 APIResponse
alibaba.ailab.user.authorized.cancel

三方用户取消授权给天猫精灵用户
*/
type AlibabaAilabUserAuthorizedCancelAPIResponse struct {
    model.CommonResponse
    AlibabaAilabUserAuthorizedCancelResponse
}

type AlibabaAilabUserAuthorizedCancelResponse struct {
    XMLName xml.Name `xml:"alibaba_ailab_user_authorized_cancel_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // message
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 返回码
    
    StatusCode   int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`

    
    // 错误中文描述
    
    Desc   string `json:"desc,omitempty" xml:"desc,omitempty"`

    
}
