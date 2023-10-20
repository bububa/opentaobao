package baichuan

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanOpenaccountRegistercodeSendAPIResponse 百川发送注册验证码 API返回值
// taobao.baichuan.openaccount.registercode.send
//
// 百川发送注册验证码
type TaobaoBaichuanOpenaccountRegistercodeSendAPIResponse struct {
	model.CommonResponse
	TaobaoBaichuanOpenaccountRegistercodeSendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBaichuanOpenaccountRegistercodeSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBaichuanOpenaccountRegistercodeSendAPIResponseModel).Reset()
}

// TaobaoBaichuanOpenaccountRegistercodeSendAPIResponseModel is 百川发送注册验证码 成功返回结果
type TaobaoBaichuanOpenaccountRegistercodeSendAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_openaccount_registercode_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBaichuanOpenaccountRegistercodeSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Name = ""
}

var poolTaobaoBaichuanOpenaccountRegistercodeSendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBaichuanOpenaccountRegistercodeSendAPIResponse)
	},
}

// GetTaobaoBaichuanOpenaccountRegistercodeSendAPIResponse 从 sync.Pool 获取 TaobaoBaichuanOpenaccountRegistercodeSendAPIResponse
func GetTaobaoBaichuanOpenaccountRegistercodeSendAPIResponse() *TaobaoBaichuanOpenaccountRegistercodeSendAPIResponse {
	return poolTaobaoBaichuanOpenaccountRegistercodeSendAPIResponse.Get().(*TaobaoBaichuanOpenaccountRegistercodeSendAPIResponse)
}

// ReleaseTaobaoBaichuanOpenaccountRegistercodeSendAPIResponse 将 TaobaoBaichuanOpenaccountRegistercodeSendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBaichuanOpenaccountRegistercodeSendAPIResponse(v *TaobaoBaichuanOpenaccountRegistercodeSendAPIResponse) {
	v.Reset()
	poolTaobaoBaichuanOpenaccountRegistercodeSendAPIResponse.Put(v)
}
