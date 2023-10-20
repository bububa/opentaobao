package bus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusCancleorderSetAPIResponse 取消订单 API返回值
// taobao.bus.cancleorder.set
//
// 取消订单
type TaobaoBusCancleorderSetAPIResponse struct {
	model.CommonResponse
	TaobaoBusCancleorderSetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBusCancleorderSetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBusCancleorderSetAPIResponseModel).Reset()
}

// TaobaoBusCancleorderSetAPIResponseModel is 取消订单 成功返回结果
type TaobaoBusCancleorderSetAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_cancleorder_set_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误代码
	ErrorCode1 string `json:"error_code1,omitempty" xml:"error_code1,omitempty"`
	// 错误描述
	ErrorMsg1 string `json:"error_msg1,omitempty" xml:"error_msg1,omitempty"`
	// success
	Success1 bool `json:"success1,omitempty" xml:"success1,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBusCancleorderSetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorCode1 = ""
	m.ErrorMsg1 = ""
	m.Success1 = false
}

var poolTaobaoBusCancleorderSetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBusCancleorderSetAPIResponse)
	},
}

// GetTaobaoBusCancleorderSetAPIResponse 从 sync.Pool 获取 TaobaoBusCancleorderSetAPIResponse
func GetTaobaoBusCancleorderSetAPIResponse() *TaobaoBusCancleorderSetAPIResponse {
	return poolTaobaoBusCancleorderSetAPIResponse.Get().(*TaobaoBusCancleorderSetAPIResponse)
}

// ReleaseTaobaoBusCancleorderSetAPIResponse 将 TaobaoBusCancleorderSetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBusCancleorderSetAPIResponse(v *TaobaoBusCancleorderSetAPIResponse) {
	v.Reset()
	poolTaobaoBusCancleorderSetAPIResponse.Put(v)
}
