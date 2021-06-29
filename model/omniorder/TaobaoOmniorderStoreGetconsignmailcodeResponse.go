package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
全渠道门店物流菜鸟裹裹取号 APIResponse
taobao.omniorder.store.getconsignmailcode

用于ISV获取全渠道门店物流订单菜鸟裹裹门店的物流快递取号
*/
type TaobaoOmniorderStoreGetconsignmailcodeAPIResponse struct {
    model.CommonResponse
    TaobaoOmniorderStoreGetconsignmailcodeResponse
}

type TaobaoOmniorderStoreGetconsignmailcodeResponse struct {
    XMLName xml.Name `xml:"omniorder_store_getconsignmailcode_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoOmniorderStoreGetconsignmailcodeResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
