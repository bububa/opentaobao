package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
打标结果查询-商品维度 APIResponse
taobao.qimen.items.tag.query

调用该接口，查询某个/某批商品上的标
*/
type TaobaoQimenItemsTagQueryAPIResponse struct {
    model.CommonResponse
    TaobaoQimenItemsTagQueryResponse
}

type TaobaoQimenItemsTagQueryResponse struct {
    XMLName xml.Name `xml:"qimen_items_tag_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // flag
    
    Flag   string `json:"flag,omitempty" xml:"flag,omitempty"`

    
    // itemTags
    
    ItemTags   []ItemTag `json:"item_tags,omitempty" xml:"item_tags>item_tag,omitempty"`
    
    
    // message
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
}
