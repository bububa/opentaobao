package idleisv

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商闲鱼商品下架 APIResponse
alibaba.idle.isv.item.downshelf

供外部服务商ISV进行闲鱼商品下架操作
*/
type AlibabaIdleIsvItemDownshelfAPIResponse struct {
    model.CommonResponse
    AlibabaIdleIsvItemDownshelfResponse
}

type AlibabaIdleIsvItemDownshelfResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_isv_item_downshelf_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果result
    
    Result   *TopResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
