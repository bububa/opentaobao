package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
打标结果查询-标维度 API返回值 
taobao.qimen.tag.items.query

调用该接口，查询打了某个标的商品列表。说明：该接口调用后，返回值的时间较长，建议不要经常调用。
*/
type TaobaoQimenTagItemsQueryAPIResponse struct {
    model.CommonResponse
    TaobaoQimenTagItemsQueryAPIResponseModel
}

// 打标结果查询-标维度 成功返回结果
type TaobaoQimenTagItemsQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"qimen_tag_items_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // flag
    Flag   string `json:"flag,omitempty" xml:"flag,omitempty"`
    // itemIds
    ItemIds   []int64 `json:"item_ids,omitempty" xml:"item_ids>int64,omitempty"`
    // message
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // tagType
    TagType   string `json:"tag_type,omitempty" xml:"tag_type,omitempty"`
}
