package baichuan

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanUserLoginAPIResponse 百川H5登录 API返回值
// taobao.baichuan.user.login
//
// 百川H5登录
type TaobaoBaichuanUserLoginAPIResponse struct {
	model.CommonResponse
	TaobaoBaichuanUserLoginAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBaichuanUserLoginAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBaichuanUserLoginAPIResponseModel).Reset()
}

// TaobaoBaichuanUserLoginAPIResponseModel is 百川H5登录 成功返回结果
type TaobaoBaichuanUserLoginAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_user_login_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBaichuanUserLoginAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Name = ""
}

var poolTaobaoBaichuanUserLoginAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBaichuanUserLoginAPIResponse)
	},
}

// GetTaobaoBaichuanUserLoginAPIResponse 从 sync.Pool 获取 TaobaoBaichuanUserLoginAPIResponse
func GetTaobaoBaichuanUserLoginAPIResponse() *TaobaoBaichuanUserLoginAPIResponse {
	return poolTaobaoBaichuanUserLoginAPIResponse.Get().(*TaobaoBaichuanUserLoginAPIResponse)
}

// ReleaseTaobaoBaichuanUserLoginAPIResponse 将 TaobaoBaichuanUserLoginAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBaichuanUserLoginAPIResponse(v *TaobaoBaichuanUserLoginAPIResponse) {
	v.Reset()
	poolTaobaoBaichuanUserLoginAPIResponse.Put(v)
}
