package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口外部订单同步 APIResponse
alibaba.wdk.order.sync

外部商户使用自助POS下单订单同步到五道口
*/
type AlibabaWdkOrderSyncAPIResponse struct {
    model.CommonResponse
    AlibabaWdkOrderSyncResponse
}

type AlibabaWdkOrderSyncResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_order_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 返回码
    
    ReturnCode   int64 `json:"return_code,omitempty" xml:"return_code,omitempty"`

    
    // 描述
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 订单号
    
    Target   string `json:"target,omitempty" xml:"target,omitempty"`

    
}
