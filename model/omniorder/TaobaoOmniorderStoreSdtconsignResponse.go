package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
通知菜鸟裹裹发货 APIResponse
taobao.omniorder.store.sdtconsign

ISV取完单号后通知菜鸟裹裹发货
*/
type TaobaoOmniorderStoreSdtconsignAPIResponse struct {
    model.CommonResponse
    TaobaoOmniorderStoreSdtconsignResponse
}

type TaobaoOmniorderStoreSdtconsignResponse struct {
    XMLName xml.Name `xml:"omniorder_store_sdtconsign_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 异常信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // data
    
    Data   *SdtConsignResponse `json:"data,omitempty" xml:"data,omitempty"`

    
    // 异常码 0 为正常，否则异常
    
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`

    
}
