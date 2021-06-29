package jms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
添加ONS消息同步用户 APIResponse
taobao.jushita.jms.user.add

添加ONS消息同步用户
*/
type TaobaoJushitaJmsUserAddAPIResponse struct {
    model.CommonResponse
    TaobaoJushitaJmsUserAddResponse
}

type TaobaoJushitaJmsUserAddResponse struct {
    XMLName xml.Name `xml:"jushita_jms_user_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功，如果失败请看错误信息
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
