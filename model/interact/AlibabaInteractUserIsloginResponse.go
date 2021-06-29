package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
校验用户是否已经登录 APIResponse
alibaba.interact.user.islogin

API的功能是校验用户是否登录，ISV调用接口的时候，通过此接口映射到mtop.interact.user.islogin接口上，因此接口只是做一个给ISV注册调用api的入口，没有发生真实的RPC
*/
type AlibabaInteractUserIsloginAPIResponse struct {
    model.CommonResponse
    AlibabaInteractUserIsloginResponse
}

type AlibabaInteractUserIsloginResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_user_islogin_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaInteractUserIsloginMtopResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
