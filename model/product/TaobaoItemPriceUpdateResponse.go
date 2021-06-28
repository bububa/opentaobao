package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新商品价格 APIResponse
taobao.item.price.update

更新商品价格
*/
type TaobaoItemPriceUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"item_price_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品结构里的num_iid，modified
    
    Item   *Item `json:"item,omitempty" xml:"