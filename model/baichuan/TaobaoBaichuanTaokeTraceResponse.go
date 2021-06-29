package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
百川淘客打点 APIResponse
taobao.baichuan.taoke.trace

百川淘客打点
*/
type TaobaoBaichuanTaokeTraceAPIResponse struct {
    model.CommonResponse
    TaobaoBaichuanTaokeTraceResponse
}

type TaobaoBaichuanTaokeTraceResponse struct {
    XMLName xml.Name `xml:"baichuan_taoke_trace_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // name
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`

    
}
