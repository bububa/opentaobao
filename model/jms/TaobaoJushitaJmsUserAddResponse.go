package jms

import (
    "github.com/bububa/opentaobao/model"
)

/* 
添加ONS消息同步用户 APIResponse
taobao.jushita.jms.user.add

添加ONS消息同步用户
*/
type TaobaoJushitaJmsUserAddAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoJushitaJmsUserAddResponse `json:"jushita_jms_user_add_response,omitempty"` 
    TaobaoJushitaJmsUserAddResponse
}

/* model for simplify = false
type TaobaoJushitaJmsUserAddResponse struct {

    // 是否成功，如果失败请看错误信息
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoJushitaJmsUserAddResponse struct {

    // 是否成功，如果失败请看错误信息
    IsSuccess   bool `json:"is_success,omitempty"`

}
