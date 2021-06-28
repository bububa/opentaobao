package openim

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量获取用户信息 APIResponse
taobao.openim.users.get

批量获取用户信息
*/
type TaobaoOpenimUsersGetAPIResponse struct {
    model.CommonResponse
    TaobaoOpenimUsersGetResponse
}

type TaobaoOpenimUsersGetResponse struct {
    XMLName xml.Name `xml:"openim_users_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 获取的用户信息列表
    
    Userinfos   []Userinfos `json:"userinfos,omitempty" xml:"userinfos>userinfos,omitempty"`
    
    
}
