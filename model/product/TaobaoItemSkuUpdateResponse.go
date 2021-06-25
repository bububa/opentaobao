package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新SKU信息 APIResponse
taobao.item.sku.update

*更新一个sku的数据 <br/>*需要更新的sku通过属性properties进行匹配查找 <br/>*商品的数量和价格必须大于等于0 <br/>*sku记录会更新到指定的num_iid对应的商品中 <br/>*num_iid对应的商品必须属于当前的会话用户
*/
type TaobaoItemSkuUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoItemSkuUpdateResponse `json:"taobao_item_sku_update_response,omitempty"`
}

type TaobaoItemSkuUpdateResponse struct {

    // 商品Sku
    Sku   *Sku `json:"sku,omitempty"`

}
