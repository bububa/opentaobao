package jms

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除ONS消息同步用户 APIResponse
taobao.jushita.jms.user.delete

删除ONS消息同步用户，删除后用户的消息将不会推送到聚石塔的ONS中
*/
type TaobaoJushitaJmsUserDeleteAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoJushitaJmsUserDeleteResponse `json:"jushita_jms_user_delete_response,omitempty"` 
    TaobaoJushitaJmsUserDeleteResponse
}

/* model for simplify = false
type TaobaoJushitaJmsUserDeleteResponse struct {

    // 是否删除成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoJushitaJmsUserDeleteResponse struct {

    // 是否删除成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
