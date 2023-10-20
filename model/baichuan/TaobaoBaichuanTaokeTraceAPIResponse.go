package baichuan

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanTaokeTraceAPIResponse 百川淘客打点 API返回值
// taobao.baichuan.taoke.trace
//
// 百川淘客打点
type TaobaoBaichuanTaokeTraceAPIResponse struct {
	model.CommonResponse
	TaobaoBaichuanTaokeTraceAPIResponseModel
}

// TaobaoBaichuanTaokeTraceAPIResponseModel is 百川淘客打点 成功返回结果
type TaobaoBaichuanTaokeTraceAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_taoke_trace_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
