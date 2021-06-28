package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新商品SKU的价格 APIResponse
taobao.item.sku.price.update

更新商品SKU的价格
*/
type TaobaoItemSkuPriceUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"item_sku_price_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品SKU信息（只包含num_iid和modified）
    
    Sku   *Sku `json:"sku,omitempty" xml:"