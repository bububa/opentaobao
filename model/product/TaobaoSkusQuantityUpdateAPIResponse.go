package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
SKU库存修改 API返回值 
taobao.skus.quantity.update

提供按照全量/增量的方式批量修改SKU库存的功能
*/
type TaobaoSkusQuantityUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoSkusQuantityUpdateAPIResponseModel
}

// SKU库存修改 成功返回结果
type TaobaoSkusQuantityUpdateAPIResponseModel struct {
    XMLName xml.Name `xml:"skus_quantity_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // iid、numIid、num和modified，skus中每个sku的skuId、quantity和modified
    Item   *Item `json:"item,omitempty" xml:"item,omitempty"`
}
