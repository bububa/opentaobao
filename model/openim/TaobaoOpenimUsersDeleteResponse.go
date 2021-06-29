package openim

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除用户 APIResponse
taobao.openim.users.delete

批量删除用户
*/
type TaobaoOpenimUsersDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoOpenimUsersDeleteResponse
}

type TaobaoOpenimUsersDeleteResponse struct {
    XMLName xml.Name `xml:"openim_users_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 操作成功
    
    Result   []string `json:"result,omitempty" xml:"result>string,omitempty"`
    
    
}
