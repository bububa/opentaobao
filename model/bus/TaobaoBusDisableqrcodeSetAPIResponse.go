package bus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusDisableqrcodeSetAPIResponse 自助机失效二维码 API返回值
// taobao.bus.disableqrcode.set
//
// 使创建的二维码失效
type TaobaoBusDisableqrcodeSetAPIResponse struct {
	model.CommonResponse
	TaobaoBusDisableqrcodeSetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBusDisableqrcodeSetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBusDisableqrcodeSetAPIResponseModel).Reset()
}

// TaobaoBusDisableqrcodeSetAPIResponseModel is 自助机失效二维码 成功返回结果
type TaobaoBusDisableqrcodeSetAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_disableqrcode_set_response"`
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
func (m *TaobaoBusDisableqrcodeSetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = ""
	m.ResultMsg = ""
	m.IsSuccess = false
}

var poolTaobaoBusDisableqrcodeSetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBusDisableqrcodeSetAPIResponse)
	},
}

// GetTaobaoBusDisableqrcodeSetAPIResponse 从 sync.Pool 获取 TaobaoBusDisableqrcodeSetAPIResponse
func GetTaobaoBusDisableqrcodeSetAPIResponse() *TaobaoBusDisableqrcodeSetAPIResponse {
	return poolTaobaoBusDisableqrcodeSetAPIResponse.Get().(*TaobaoBusDisableqrcodeSetAPIResponse)
}

// ReleaseTaobaoBusDisableqrcodeSetAPIResponse 将 TaobaoBusDisableqrcodeSetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBusDisableqrcodeSetAPIResponse(v *TaobaoBusDisableqrcodeSetAPIResponse) {
	v.Reset()
	poolTaobaoBusDisableqrcodeSetAPIResponse.Put(v)
}
