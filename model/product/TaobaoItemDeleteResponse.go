package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除单条商品 APIResponse
taobao.item.delete

删除单条商品
*/
type TaobaoItemDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoItemDeleteResponse
}

type TaobaoItemDeleteResponse struct {
    XMLName xml.Name `xml:"item_delete_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 被删除商品的相关信息
    
    Item   *Item `json:"item,omitempty" xml:"item,omitempty"`

    
}
