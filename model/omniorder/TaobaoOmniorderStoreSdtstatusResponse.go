package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟裹裹运单状态查询 APIResponse
taobao.omniorder.store.sdtstatus

提供给商家查询运力单的状态。
*/
type TaobaoOmniorderStoreSdtstatusAPIResponse struct {
    model.CommonResponse
    TaobaoOmniorderStoreSdtstatusResponse
}

type TaobaoOmniorderStoreSdtstatusResponse struct {
    XMLName xml.Name `xml:"omniorder_store_sdtstatus_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoOmniorderStoreSdtstatusResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
