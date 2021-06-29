package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单回流接口 APIResponse
alibaba.alsc.crm.open.order.backflow

回流isv订单接口
*/
type AlibabaAlscCrmOpenOrderBackflowAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmOpenOrderBackflowResponse
}

type AlibabaAlscCrmOpenOrderBackflowResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_open_order_backflow_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
