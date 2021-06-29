package idleisv

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商闲鱼商品查询 APIResponse
alibaba.idle.isv.item.query

服务商ISV闲鱼商品查询
*/
type AlibabaIdleIsvItemQueryAPIResponse struct {
    model.CommonResponse
    AlibabaIdleIsvItemQueryResponse
}

type AlibabaIdleIsvItemQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_isv_item_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果result
    
    Result   *TopResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
