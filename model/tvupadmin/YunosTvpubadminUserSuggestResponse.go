package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取关联账户列表 APIResponse
yunos.tvpubadmin.user.suggest

获取关联账户列表
*/
type YunosTvpubadminUserSuggestAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminUserSuggestResponse
}

type YunosTvpubadminUserSuggestResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_user_suggest_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // object
    
    List   []AccountSuggestDO `json:"list,omitempty" xml:"list>account_suggest_do,omitempty"`
    
    
}
