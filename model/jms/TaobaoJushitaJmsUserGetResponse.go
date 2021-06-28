package jms

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询某个用户是否同步消息 APIResponse
taobao.jushita.jms.user.get

查询某个用户是否同步消息，只支持单个查询
*/
type TaobaoJushitaJmsUserGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoJushitaJmsUserGetResponse `json:"jushita_jms_user_get_response,omitempty"` 
    TaobaoJushitaJmsUserGetResponse
}

/* model for simplify = false
type TaobaoJushitaJmsUserGetResponse struct {

    // 同步的用户信息
    
    OnsUser  *struct {
        TmcUser  *TmcUser `json:"tmc_user,omitempty"`
    } `json:"ons_user,omitempty"`
    

}
*/

type TaobaoJushitaJmsUserGetResponse struct {

    // 同步的用户信息
    OnsUser   *TmcUser `json:"ons_user,omitempty"`

}
