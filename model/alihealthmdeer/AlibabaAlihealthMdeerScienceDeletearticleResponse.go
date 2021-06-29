package alihealthmdeer

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
文章删除 APIResponse
alibaba.alihealth.mdeer.science.deletearticle

三方同步文章删除
*/
type AlibabaAlihealthMdeerScienceDeletearticleAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthMdeerScienceDeletearticleResponse
}

type AlibabaAlihealthMdeerScienceDeletearticleResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_mdeer_science_deletearticle_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 错误信息
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
    // 错误code
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
    // 是否删除成功
    
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`

    
}
