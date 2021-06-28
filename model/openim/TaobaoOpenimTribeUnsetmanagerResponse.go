package openim

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群取消管理员 APIResponse
taobao.openim.tribe.unsetmanager

OPENIM群取消管理员
*/
type TaobaoOpenimTribeUnsetmanagerAPIResponse struct {
    model.CommonResponse
    TaobaoOpenimTribeUnsetmanagerResponse
}

type TaobaoOpenimTribeUnsetmanagerResponse struct {
    XMLName xml.Name `xml:"openim_tribe_unsetmanager_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 群服务code
    
    TribeCode   int64 `json:"tribe_code,omitempty" xml:"tribe_code,omitempty"`

    
}
