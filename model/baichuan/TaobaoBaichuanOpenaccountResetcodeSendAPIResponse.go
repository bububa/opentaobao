package baichuan

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanOpenaccountResetcodeSendAPIResponse 百川发送找回密码验证码 API返回值
// taobao.baichuan.openaccount.resetcode.send
//
// 百川发送找回密码验证码
type TaobaoBaichuanOpenaccountResetcodeSendAPIResponse struct {
	model.CommonResponse
	TaobaoBaichuanOpenaccountResetcodeSendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBaichuanOpenaccountResetcodeSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBaichuanOpenaccountResetcodeSendAPIResponseModel).Reset()
}

// TaobaoBaichuanOpenaccountResetcodeSendAPIResponseModel is 百川发送找回密码验证码 成功返回结果
type TaobaoBaichuanOpenaccountResetcodeSendAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_openaccount_resetcode_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBaichuanOpenaccountResetcodeSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Name = ""
}

var poolTaobaoBaichuanOpenaccountResetcodeSendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBaichuanOpenaccountResetcodeSendAPIResponse)
	},
}

// GetTaobaoBaichuanOpenaccountResetcodeSendAPIResponse 从 sync.Pool 获取 TaobaoBaichuanOpenaccountResetcodeSendAPIResponse
func GetTaobaoBaichuanOpenaccountResetcodeSendAPIResponse() *TaobaoBaichuanOpenaccountResetcodeSendAPIResponse {
	return poolTaobaoBaichuanOpenaccountResetcodeSendAPIResponse.Get().(*TaobaoBaichuanOpenaccountResetcodeSendAPIResponse)
}

// ReleaseTaobaoBaichuanOpenaccountResetcodeSendAPIResponse 将 TaobaoBaichuanOpenaccountResetcodeSendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBaichuanOpenaccountResetcodeSendAPIResponse(v *TaobaoBaichuanOpenaccountResetcodeSendAPIResponse) {
	v.Reset()
	poolTaobaoBaichuanOpenaccountResetcodeSendAPIResponse.Put(v)
}
