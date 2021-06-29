package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
租赁商品发布 APIResponse
alibaba.idle.rent.item.add

发布闲鱼租赁商品
*/
type AlibabaIdleRentItemAddAPIResponse struct {
    model.CommonResponse
    AlibabaIdleRentItemAddResponse
}

type AlibabaIdleRentItemAddResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_rent_item_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 系统自动生成
    
    Result   *TopResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
