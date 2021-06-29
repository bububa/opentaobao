package fans

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
消息推送 APIResponse
tmall.fans.arena.push

超级擂台消息推送
*/
type TmallFansArenaPushAPIResponse struct {
    model.CommonResponse
    TmallFansArenaPushResponse
}

type TmallFansArenaPushResponse struct {
    XMLName xml.Name `xml:"tmall_fans_arena_push_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回对象
    
    Result   *FansResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
