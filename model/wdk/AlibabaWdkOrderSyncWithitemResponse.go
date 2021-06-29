package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单和商品同步接口 APIResponse
alibaba.wdk.order.sync.withitem

轻pos,将订单和商品的信息一起传到盒马这边，进行创单和添加商品处理。
*/
type AlibabaWdkOrderSyncWithitemAPIResponse struct {
    model.CommonResponse
    AlibabaWdkOrderSyncWithitemResponse
}

type AlibabaWdkOrderSyncWithitemResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_order_sync_withitem_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 调用返回的结果
    
    Result   *AlibabaWdkOrderSyncWithitemApiResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
