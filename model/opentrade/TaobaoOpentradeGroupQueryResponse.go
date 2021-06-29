package opentrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
组团购场景查询商品组团详情 APIResponse
taobao.opentrade.group.query

组团购场景下，查询商品开团详情
*/
type TaobaoOpentradeGroupQueryAPIResponse struct {
    model.CommonResponse
    TaobaoOpentradeGroupQueryResponse
}

type TaobaoOpentradeGroupQueryResponse struct {
    XMLName xml.Name `xml:"opentrade_group_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 总记录数
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`

    
    // 团信息
    
    Results   []GroupDetailResponse `json:"results,omitempty" xml:"results>group_detail_response,omitempty"`
    
    
}
