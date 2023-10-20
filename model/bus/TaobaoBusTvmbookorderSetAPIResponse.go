package bus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusTvmbookorderSetAPIResponse 线下自助机通知出票接口 API返回值
// taobao.bus.tvmbookorder.set
//
// 出票，当成功的时候告知出票；当失败的时候告知出票失败，飞猪退款给用户。
type TaobaoBusTvmbookorderSetAPIResponse struct {
	model.CommonResponse
	TaobaoBusTvmbookorderSetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBusTvmbookorderSetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBusTvmbookorderSetAPIResponseModel).Reset()
}

// TaobaoBusTvmbookorderSetAPIResponseModel is 线下自助机通知出票接口 成功返回结果
type TaobaoBusTvmbookorderSetAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_tvmbookorder_set_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// errorCode
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// errorMsg
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBusTvmbookorderSetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = ""
	m.ResultMsg = ""
	m.IsSuccess = false
}

var poolTaobaoBusTvmbookorderSetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBusTvmbookorderSetAPIResponse)
	},
}

// GetTaobaoBusTvmbookorderSetAPIResponse 从 sync.Pool 获取 TaobaoBusTvmbookorderSetAPIResponse
func GetTaobaoBusTvmbookorderSetAPIResponse() *TaobaoBusTvmbookorderSetAPIResponse {
	return poolTaobaoBusTvmbookorderSetAPIResponse.Get().(*TaobaoBusTvmbookorderSetAPIResponse)
}

// ReleaseTaobaoBusTvmbookorderSetAPIResponse 将 TaobaoBusTvmbookorderSetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBusTvmbookorderSetAPIResponse(v *TaobaoBusTvmbookorderSetAPIResponse) {
	v.Reset()
	poolTaobaoBusTvmbookorderSetAPIResponse.Put(v)
}
