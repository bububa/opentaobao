package opentrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
交易开放获取商品绑定信息 APIResponse
taobao.opentrade.tools.items.query

交易开放获取商品绑定信息
*/
type TaobaoOpentradeToolsItemsQueryAPIResponse struct {
    model.CommonResponse
    TaobaoOpentradeToolsItemsQueryResponse
}

type TaobaoOpentradeToolsItemsQueryResponse struct {
    XMLName xml.Name `xml:"opentrade_tools_items_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 已绑定的商品ID列表
    
    ItemIds   []int64 `json:"item_ids,omitempty" xml:"item_ids>int64,omitempty"`
    
    
}
