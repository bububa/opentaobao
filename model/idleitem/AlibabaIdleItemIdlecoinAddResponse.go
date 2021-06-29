package idleitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
免费送商品发送 APIResponse
alibaba.idle.item.idlecoin.add

免费送商品发布
*/
type AlibabaIdleItemIdlecoinAddAPIResponse struct {
    model.CommonResponse
    AlibabaIdleItemIdlecoinAddResponse
}

type AlibabaIdleItemIdlecoinAddResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_item_idlecoin_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *IdleResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
