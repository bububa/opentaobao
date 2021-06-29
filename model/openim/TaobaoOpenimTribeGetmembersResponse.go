package openim

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群成员获取 APIResponse
taobao.openim.tribe.getmembers

OPENIM群成员获取
*/
type TaobaoOpenimTribeGetmembersAPIResponse struct {
    model.CommonResponse
    TaobaoOpenimTribeGetmembersResponse
}

type TaobaoOpenimTribeGetmembersResponse struct {
    XMLName xml.Name `xml:"openim_tribe_getmembers_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // OPENIM群成员列表
    
    TribeUserList   []TribeUser `json:"tribe_user_list,omitempty" xml:"tribe_user_list>tribe_user,omitempty"`
    
    
}
