package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
通知速店通取消取号 APIResponse
taobao.omniorder.store.sdtcancel

通知速店通取消取号
*/
type TaobaoOmniorderStoreSdtcancelAPIResponse struct {
    model.CommonResponse
    TaobaoOmniorderStoreSdtcancelResponse
}

type TaobaoOmniorderStoreSdtcancelResponse struct {
    XMLName xml.Name `xml:"omniorder_store_sdtcancel_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TaobaoOmniorderStoreSdtcancelResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
