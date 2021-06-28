package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取当前授权用户手机号码 APIResponse
taobao.miniapp.user.phone.get

在商家应用中，获取当前授权用户手机号码
*/
type TaobaoMiniappUserPhoneGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoMiniappUserPhoneGetResponse `json:"miniapp_user_phone_get_response,omitempty"` 
    TaobaoMiniappUserPhoneGetResponse
}

/* model for simplify = false
type TaobaoMiniappUserPhoneGetResponse struct {

    // 用户手机号码
    
    Phone   string `json:"phone,omitempty"`
    

}
*/

type TaobaoMiniappUserPhoneGetResponse struct {

    // 用户手机号码
    Phone   string `json:"phone,omitempty"`

}
