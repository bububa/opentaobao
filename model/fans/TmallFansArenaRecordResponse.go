package fans

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
记录完成擂台的用户 APIResponse
tmall.fans.arena.record

记录完成擂台的用户和完成分数
*/
type TmallFansArenaRecordAPIResponse struct {
    model.CommonResponse
    TmallFansArenaRecordResponse
}

type TmallFansArenaRecordResponse struct {
    XMLName xml.Name `xml:"tmall_fans_arena_record_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回对象
    
    FansResult   *FansResult `json:"fans_result,omitempty" xml:"fans_result,omitempty"`

    
}
