package openim

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群信息修改 APIResponse
taobao.openim.tribe.modifytribeinfo

OPENIM群信息修改
*/
type TaobaoOpenimTribeModifytribeinfoAPIResponse struct {
    model.CommonResponse
    TaobaoOpenimTribeModifytribeinfoResponse
}

type TaobaoOpenimTribeModifytribeinfoResponse struct {
    XMLName xml.Name `xml:"openim_tribe_modifytribeinfo_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 群服务code
    
    TribeCode   int64 `json:"tribe_code,omitempty" xml:"tribe_code,omitempty"`

    
}
