package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除SKU APIResponse
taobao.item.sku.delete

删除一个sku的数据<br/>需要删除的sku通过属性properties进行匹配查找
*/
type TaobaoItemSkuDeleteAPIResponse struct {
    model.CommonResponse
    Response *TaobaoItemSkuDeleteResponse `json:"taobao_item_sku_delete_response,omitempty"`
}

type TaobaoItemSkuDeleteResponse struct {

    // Sku结构
    Sku   *Sku `json:"sku,omitempty"`

}
