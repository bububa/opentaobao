package tmallfcbox

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
丰巢通知接口 APIResponse
tmall.fcbox.notify

tmax接收丰巢快递通知
*/
type TmallFcboxNotifyAPIResponse struct {
    model.CommonResponse
    TmallFcboxNotifyResponse
}

type TmallFcboxNotifyResponse struct {
    XMLName xml.Name `xml:"tmall_fcbox_notify_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TmallFcboxNotifyResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
