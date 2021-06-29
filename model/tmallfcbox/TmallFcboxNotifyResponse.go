package tmallfcbox

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
丰巢通知接口 API返回值 
tmall.fcbox.notify

tmax接收丰巢快递通知
*/
type TmallFcboxNotifyAPIResponse struct {
    model.CommonResponse
    TmallFcboxNotifyResponse
}

// 丰巢通知接口 成功返回结果
type TmallFcboxNotifyResponse struct {
    XMLName xml.Name `xml:"tmall_fcbox_notify_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TmallFcboxNotifyResult `json:"result,omitempty" xml:"result,omitempty"`
}
