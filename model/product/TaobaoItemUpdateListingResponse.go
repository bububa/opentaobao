package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
一口价商品上架 APIResponse
taobao.item.update.listing

* 单个商品上架<br/>* 输入的num_iid必须属于当前会话用户
*/
type TaobaoItemUpdateListingAPIResponse struct {
    model.CommonResponse
    TaobaoItemUpdateListingResponse
}

type TaobaoItemUpdateListingResponse struct {
    XMLName xml.Name `xml:"item_update_listing_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 上架后返回的商品信息：返回的结果就是:num_iid和modified
    
    Item   *Item `json:"item,omitempty" xml:"item,omitempty"`

    
}
