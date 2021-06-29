package rhino

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
WMS003成衣入库确认 APIResponse
taobao.rhino.supplychain.inbound.confirm

【WMS003】【同步成衣入库完成信息】
*/
type TaobaoRhinoSupplychainInboundConfirmAPIResponse struct {
    model.CommonResponse
    TaobaoRhinoSupplychainInboundConfirmResponse
}

type TaobaoRhinoSupplychainInboundConfirmResponse struct {
    XMLName xml.Name `xml:"rhino_supplychain_inbound_confirm_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // message
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // code
    
    RetCode   int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`

    
    // success
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
