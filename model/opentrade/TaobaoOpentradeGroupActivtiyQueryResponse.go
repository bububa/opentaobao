package opentrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
组团活动信息查询 APIResponse
taobao.opentrade.group.activtiy.query

组团购场景下，团购活动信息查询
*/
type TaobaoOpentradeGroupActivtiyQueryAPIResponse struct {
    model.CommonResponse
    TaobaoOpentradeGroupActivtiyQueryResponse
}

type TaobaoOpentradeGroupActivtiyQueryResponse struct {
    XMLName xml.Name `xml:"opentrade_group_activtiy_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 总记录数
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`

    
    // 组团活动信息
    
    Results   []GroupActivityResponse `json:"results,omitempty" xml:"results>group_activity_response,omitempty"`
    
    
}
