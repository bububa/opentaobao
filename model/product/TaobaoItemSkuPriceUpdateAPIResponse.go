package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新商品SKU的价格 API返回值 
taobao.item.sku.price.update

更新商品SKU的价格
*/
type TaobaoItemSkuPriceUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoItemSkuPriceUpdateAPIResponseModel
}

// 更新商品SKU的价格 成功返回结果
type TaobaoItemSkuPriceUpdateAPIResponseModel struct {
    XMLName xml.Name `xml:"item_sku_price_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 商品SKU信息（只包含num_iid和modified）
    Sku   *Sku `json:"sku,omitempty" xml:"sku,omitempty"`
}
