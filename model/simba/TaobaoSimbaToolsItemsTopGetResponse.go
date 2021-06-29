package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取得一个关键词的推广组排名列表 APIResponse
taobao.simba.tools.items.top.get

取得一个关键词的推广组排名列表
*/
type TaobaoSimbaToolsItemsTopGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaToolsItemsTopGetResponse
}

type TaobaoSimbaToolsItemsTopGetResponse struct {
    XMLName xml.Name `xml:"simba_tools_items_top_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 推广组信息列表
    
    Rankeditems   []RankedItem `json:"rankeditems,omitempty" xml:"rankeditems>ranked_item,omitempty"`
    
    
}
