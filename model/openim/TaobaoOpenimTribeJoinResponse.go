package openim

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群主动加入 APIResponse
taobao.openim.tribe.join

OPENIM群主动加入
*/
type TaobaoOpenimTribeJoinAPIResponse struct {
    model.CommonResponse
    TaobaoOpenimTribeJoinResponse
}

type TaobaoOpenimTribeJoinResponse struct {
    XMLName xml.Name `xml:"openim_tribe_join_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 群服务code
    
    TribeCode   int64 `json:"tribe_code,omitempty" xml:"tribe_code,omitempty"`

    
}
