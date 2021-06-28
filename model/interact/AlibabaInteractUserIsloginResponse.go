package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
校验用户是否已经登录 APIResponse
alibaba.interact.user.islogin

API的功能是校验用户是否登录，ISV调用接口的时候，通过此接口映射到mtop.interact.user.islogin接口上，因此接口只是做一个给ISV注册调用api的入口，没有发生真实的RPC
*/
type AlibabaInteractUserIsloginAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaInteractUserIsloginResponse `json:"alibaba_interact_user_islogin_response,omitempty"` 
    AlibabaInteractUserIsloginResponse
}

/* model for simplify = false
type AlibabaInteractUserIsloginResponse struct {

    // result
    
    Result  *struct {
        MtopResult  *MtopResult `json:"mtop_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaInteractUserIsloginResponse struct {

    // result
    Result   *MtopResult `json:"result,omitempty"`

}
