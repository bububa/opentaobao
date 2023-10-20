package bus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusNumbersUpdateAPIResponse 汽车票车次更新服务 API返回值
// taobao.bus.numbers.update
//
// 用于汽车票车次信息的新增、更新和逻辑删除
type TaobaoBusNumbersUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoBusNumbersUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBusNumbersUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBusNumbersUpdateAPIResponseModel).Reset()
}

// TaobaoBusNumbersUpdateAPIResponseModel is 汽车票车次更新服务 成功返回结果
type TaobaoBusNumbersUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_numbers_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 成功数量
	SuccCount int64 `json:"succ_count,omitempty" xml:"succ_count,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBusNumbersUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = ""
	m.ResultMsg = ""
	m.SuccCount = 0
	m.IsSuccess = false
}

var poolTaobaoBusNumbersUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBusNumbersUpdateAPIResponse)
	},
}

// GetTaobaoBusNumbersUpdateAPIResponse 从 sync.Pool 获取 TaobaoBusNumbersUpdateAPIResponse
func GetTaobaoBusNumbersUpdateAPIResponse() *TaobaoBusNumbersUpdateAPIResponse {
	return poolTaobaoBusNumbersUpdateAPIResponse.Get().(*TaobaoBusNumbersUpdateAPIResponse)
}

// ReleaseTaobaoBusNumbersUpdateAPIResponse 将 TaobaoBusNumbersUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBusNumbersUpdateAPIResponse(v *TaobaoBusNumbersUpdateAPIResponse) {
	v.Reset()
	poolTaobaoBusNumbersUpdateAPIResponse.Put(v)
}
