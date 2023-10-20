package baichuan

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanUserLoginbytokenAPIResponse 百川手淘信任登录 API返回值
// taobao.baichuan.user.loginbytoken
//
// 百川手淘信任登录
type TaobaoBaichuanUserLoginbytokenAPIResponse struct {
	model.CommonResponse
	TaobaoBaichuanUserLoginbytokenAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBaichuanUserLoginbytokenAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBaichuanUserLoginbytokenAPIResponseModel).Reset()
}

// TaobaoBaichuanUserLoginbytokenAPIResponseModel is 百川手淘信任登录 成功返回结果
type TaobaoBaichuanUserLoginbytokenAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_user_loginbytoken_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBaichuanUserLoginbytokenAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Name = ""
}

var poolTaobaoBaichuanUserLoginbytokenAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBaichuanUserLoginbytokenAPIResponse)
	},
}

// GetTaobaoBaichuanUserLoginbytokenAPIResponse 从 sync.Pool 获取 TaobaoBaichuanUserLoginbytokenAPIResponse
func GetTaobaoBaichuanUserLoginbytokenAPIResponse() *TaobaoBaichuanUserLoginbytokenAPIResponse {
	return poolTaobaoBaichuanUserLoginbytokenAPIResponse.Get().(*TaobaoBaichuanUserLoginbytokenAPIResponse)
}

// ReleaseTaobaoBaichuanUserLoginbytokenAPIResponse 将 TaobaoBaichuanUserLoginbytokenAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBaichuanUserLoginbytokenAPIResponse(v *TaobaoBaichuanUserLoginbytokenAPIResponse) {
	v.Reset()
	poolTaobaoBaichuanUserLoginbytokenAPIResponse.Put(v)
}
