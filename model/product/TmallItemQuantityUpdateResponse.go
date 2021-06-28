package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫商品/SKU库存更新接口 APIResponse
tmall.item.quantity.update

天猫商品/SKU库存更新接口；支持商品库存更新；支持同一商品下的SKU批量更新。
*/
type TmallItemQuantityUpdateAPIResponse struct {
    model.CommonResponse
    TmallItemQuantityUpdateResponse
}

type TmallItemQuantityUpdateResponse struct {
    XMLName xml.Name `xml:"tmall_item_quantity_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 库存更新结果，商品id
    
    QuantityUpdateResult   string `json:"quantity_update_result,omitempty" xml:"quantity_update_result,omitempty"`

    
}
