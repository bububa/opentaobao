package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
租赁商品编辑 APIResponse
alibaba.idle.rent.item.edit

发布闲鱼租赁商品
*/
type AlibabaIdleRentItemEditAPIResponse struct {
    model.CommonResponse
    AlibabaIdleRentItemEditResponse
}

type AlibabaIdleRentItemEditResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_rent_item_edit_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 系统自动生成
    
    Result   *TopResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
