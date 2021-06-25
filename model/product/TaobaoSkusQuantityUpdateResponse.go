package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
SKU库存修改 APIResponse
taobao.skus.quantity.update

提供按照全量/增量的方式批量修改SKU库存的功能
*/
type TaobaoSkusQuantityUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSkusQuantityUpdateResponse `json:"taobao_skus_quantity_update_response,omitempty"`
}

type TaobaoSkusQuantityUpdateResponse struct {

    // iid、numIid、num和modified，skus中每个sku的skuId、quantity和modified
    Item   *Item `json:"item,omitempty"`

}
