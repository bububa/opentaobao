package icbuproduct

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
icbu商品库存更新 APIResponse
alibaba.icbu.product.inventory.update

更新库存信息
*/
type AlibabaIcbuProductInventoryUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaIcbuProductInventoryUpdateResponse
}

type AlibabaIcbuProductInventoryUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_icbu_product_inventory_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // Top返回对象
    
    Result   *TopResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
