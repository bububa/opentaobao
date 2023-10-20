package baichuan

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanOpenaccountResetcodeCheckAPIResponse 百川验证找回密码验证码 API返回值
// taobao.baichuan.openaccount.resetcode.check
//
// 百川验证找回密码验证码
type TaobaoBaichuanOpenaccountResetcodeCheckAPIResponse struct {
	model.CommonResponse
	TaobaoBaichuanOpenaccountResetcodeCheckAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBaichuanOpenaccountResetcodeCheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBaichuanOpenaccountResetcodeCheckAPIResponseModel).Reset()
}

// TaobaoBaichuanOpenaccountResetcodeCheckAPIResponseModel is 百川验证找回密码验证码 成功返回结果
type TaobaoBaichuanOpenaccountResetcodeCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_openaccount_resetcode_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBaichuanOpenaccountResetcodeCheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Name = ""
}

var poolTaobaoBaichuanOpenaccountResetcodeCheckAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBaichuanOpenaccountResetcodeCheckAPIResponse)
	},
}

// GetTaobaoBaichuanOpenaccountResetcodeCheckAPIResponse 从 sync.Pool 获取 TaobaoBaichuanOpenaccountResetcodeCheckAPIResponse
func GetTaobaoBaichuanOpenaccountResetcodeCheckAPIResponse() *TaobaoBaichuanOpenaccountResetcodeCheckAPIResponse {
	return poolTaobaoBaichuanOpenaccountResetcodeCheckAPIResponse.Get().(*TaobaoBaichuanOpenaccountResetcodeCheckAPIResponse)
}

// ReleaseTaobaoBaichuanOpenaccountResetcodeCheckAPIResponse 将 TaobaoBaichuanOpenaccountResetcodeCheckAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBaichuanOpenaccountResetcodeCheckAPIResponse(v *TaobaoBaichuanOpenaccountResetcodeCheckAPIResponse) {
	v.Reset()
	poolTaobaoBaichuanOpenaccountResetcodeCheckAPIResponse.Put(v)
}
