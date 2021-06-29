package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新/增加sku信息 APIResponse
alibaba.idle.rent.item.sku.update

更新/增加sku信息
*/
type AlibabaIdleRentItemSkuUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaIdleRentItemSkuUpdateResponse
}

type AlibabaIdleRentItemSkuUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_rent_item_sku_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 系统自动生成
    
    Result   *TopResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
