package baichuan

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoBaichuanTaokeTraceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBaichuanTaokeTraceAPIResponseModel).Reset()
}

// TaobaoBaichuanTaokeTraceAPIResponseModel is 百川淘客打点 成功返回结果
type TaobaoBaichuanTaokeTraceAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_taoke_trace_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBaichuanTaokeTraceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Name = ""
}

var poolTaobaoBaichuanTaokeTraceAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBaichuanTaokeTraceAPIResponse)
	},
}

// GetTaobaoBaichuanTaokeTraceAPIResponse 从 sync.Pool 获取 TaobaoBaichuanTaokeTraceAPIResponse
func GetTaobaoBaichuanTaokeTraceAPIResponse() *TaobaoBaichuanTaokeTraceAPIResponse {
	return poolTaobaoBaichuanTaokeTraceAPIResponse.Get().(*TaobaoBaichuanTaokeTraceAPIResponse)
}

// ReleaseTaobaoBaichuanTaokeTraceAPIResponse 将 TaobaoBaichuanTaokeTraceAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBaichuanTaokeTraceAPIResponse(v *TaobaoBaichuanTaokeTraceAPIResponse) {
	v.Reset()
	poolTaobaoBaichuanTaokeTraceAPIResponse.Put(v)
}
