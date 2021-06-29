package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
添加数据推送用户 APIResponse
taobao.jushita.jdp.user.add

提供给接入数据推送的应用添加数据推送服务的用户
*/
type TaobaoJushitaJdpUserAddAPIResponse struct {
    model.CommonResponse
    TaobaoJushitaJdpUserAddResponse
}

type TaobaoJushitaJdpUserAddResponse struct {
    XMLName xml.Name `xml:"jushita_jdp_user_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否添加成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
