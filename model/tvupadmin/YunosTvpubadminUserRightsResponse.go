package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户权益 APIResponse
yunos.tvpubadmin.user.rights

获取用户权益
*/
type YunosTvpubadminUserRightsAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminUserRightsResponse
}

type YunosTvpubadminUserRightsResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_user_rights_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // object
    
    Object   *PaginationDO `json:"object,omitempty" xml:"object,omitempty"`

    
}
