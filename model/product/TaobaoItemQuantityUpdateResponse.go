package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
宝贝/SKU库存修改 APIResponse
taobao.item.quantity.update

提供按照全量或增量形式修改宝贝/SKU库存的功能
*/
type TaobaoItemQuantityUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoItemQuantityUpdateResponse
}

type TaobaoItemQuantityUpdateResponse struct {
    XMLName xml.Name `xml:"item_quantity_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // iid、numIid、num和modified，skus中每个sku的skuId、quantity和modified
    
    Item   *Item `json:"item,omitempty" xml:"item,omitempty"`

    
}
