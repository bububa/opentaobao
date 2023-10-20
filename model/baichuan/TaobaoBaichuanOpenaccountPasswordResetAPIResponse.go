package baichuan

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanOpenaccountPasswordResetAPIResponse 百川找回密码 API返回值
// taobao.baichuan.openaccount.password.reset
//
// 百川找回密码
type TaobaoBaichuanOpenaccountPasswordResetAPIResponse struct {
	model.CommonResponse
	TaobaoBaichuanOpenaccountPasswordResetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBaichuanOpenaccountPasswordResetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBaichuanOpenaccountPasswordResetAPIResponseModel).Reset()
}

// TaobaoBaichuanOpenaccountPasswordResetAPIResponseModel is 百川找回密码 成功返回结果
type TaobaoBaichuanOpenaccountPasswordResetAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_openaccount_password_reset_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBaichuanOpenaccountPasswordResetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Name = ""
}

var poolTaobaoBaichuanOpenaccountPasswordResetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBaichuanOpenaccountPasswordResetAPIResponse)
	},
}

// GetTaobaoBaichuanOpenaccountPasswordResetAPIResponse 从 sync.Pool 获取 TaobaoBaichuanOpenaccountPasswordResetAPIResponse
func GetTaobaoBaichuanOpenaccountPasswordResetAPIResponse() *TaobaoBaichuanOpenaccountPasswordResetAPIResponse {
	return poolTaobaoBaichuanOpenaccountPasswordResetAPIResponse.Get().(*TaobaoBaichuanOpenaccountPasswordResetAPIResponse)
}

// ReleaseTaobaoBaichuanOpenaccountPasswordResetAPIResponse 将 TaobaoBaichuanOpenaccountPasswordResetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBaichuanOpenaccountPasswordResetAPIResponse(v *TaobaoBaichuanOpenaccountPasswordResetAPIResponse) {
	v.Reset()
	poolTaobaoBaichuanOpenaccountPasswordResetAPIResponse.Put(v)
}
