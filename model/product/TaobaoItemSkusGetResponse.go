package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据商品ID列表获取SKU信息 APIResponse
taobao.item.skus.get

* 获取多个商品下的所有sku
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong>
*/
type TaobaoItemSkusGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoItemSkusGetResponse `json:"taobao_item_skus_get_response,omitempty"`
}

type TaobaoItemSkusGetResponse struct {

    // Sku列表
    Skus   []Sku `json:"skus,omitempty"`

}
