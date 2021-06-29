package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分佣数据传输 APIResponse
alibaba.retail.commission.order.sync

同步分佣结果
*/
type AlibabaRetailCommissionOrderSyncAPIResponse struct {
    model.CommonResponse
    AlibabaRetailCommissionOrderSyncResponse
}

type AlibabaRetailCommissionOrderSyncResponse struct {
    XMLName xml.Name `xml:"alibaba_retail_commission_order_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回描述
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 返回的数据实体
    
    Data   bool `json:"data,omitempty" xml:"data,omitempty"`

    
    // 返回的执行状态吗
    
    SCode   string `json:"s_code,omitempty" xml:"s_code,omitempty"`

    
    // 是否执行成功
    
    SSuccess   bool `json:"s_success,omitempty" xml:"s_success,omitempty"`

    
}
