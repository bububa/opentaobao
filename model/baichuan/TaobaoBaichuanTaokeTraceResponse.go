package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
百川淘客打点 API返回值 
taobao.baichuan.taoke.trace

百川淘客打点
*/
type TaobaoBaichuanTaokeTraceAPIResponse struct {
    model.CommonResponse
    TaobaoBaichuanTaokeTraceResponse
}

// 百川淘客打点 成功返回结果
type TaobaoBaichuanTaokeTraceResponse struct {
    XMLName xml.Name `xml:"baichuan_taoke_trace_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // name
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
}
