package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
SKU库存修改 APIResponse
taobao.skus.quantity.update

提供按照全量/增量的方式批量修改SKU库存的功能
*/
type TaobaoSkusQuantityUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"skus_quantity_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // iid、numIid、num和modified，skus中每个sku的skuId、quantity和modified
    
    Item   *Item `json:"item,omitempty" xml:"