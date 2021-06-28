package jms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询某个用户是否同步消息 APIResponse
taobao.jushita.jms.user.get

查询某个用户是否同步消息，只支持单个查询
*/
type TaobaoJushitaJmsUserGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"jushita_jms_user_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 同步的用户信息
    
    OnsUser   *TmcUser `json:"ons_user,omitempty" xml:"