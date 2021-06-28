package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除数据推送用户 APIResponse
taobao.jushita.jdp.user.delete

删除应用的数据推送用户，用户被删除后，重新添加时会重新同步历史数据。
*/
type TaobaoJushitaJdpUserDeleteAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoJushitaJdpUserDeleteResponse `json:"jushita_jdp_user_delete_response,omitempty"` 
    TaobaoJushitaJdpUserDeleteResponse
}

/* model for simplify = false
type TaobaoJushitaJdpUserDeleteResponse struct {

    // 是否删除成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoJushitaJdpUserDeleteResponse struct {

    // 是否删除成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
