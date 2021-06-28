package jms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除ONS消息同步用户 APIResponse
taobao.jushita.jms.user.delete

删除ONS消息同步用户，删除后用户的消息将不会推送到聚石塔的ONS中
*/
type TaobaoJushitaJmsUserDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoJushitaJmsUserDeleteResponse
}

type TaobaoJushitaJmsUserDeleteResponse struct {
    XMLName xml.Name `xml:"jushita_jms_user_delete_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否删除成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
