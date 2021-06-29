package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除数据推送用户 APIResponse
taobao.jushita.jdp.user.delete

删除应用的数据推送用户，用户被删除后，重新添加时会重新同步历史数据。
*/
type TaobaoJushitaJdpUserDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoJushitaJdpUserDeleteResponse
}

type TaobaoJushitaJdpUserDeleteResponse struct {
    XMLName xml.Name `xml:"jushita_jdp_user_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否删除成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
