package baichuan

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanOpenaccountLoginAPIResponse 百川用户名密码登录 API返回值
// taobao.baichuan.openaccount.login
//
// 百川用户名密码登录
type TaobaoBaichuanOpenaccountLoginAPIResponse struct {
	model.CommonResponse
	TaobaoBaichuanOpenaccountLoginAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBaichuanOpenaccountLoginAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBaichuanOpenaccountLoginAPIResponseModel).Reset()
}

// TaobaoBaichuanOpenaccountLoginAPIResponseModel is 百川用户名密码登录 成功返回结果
type TaobaoBaichuanOpenaccountLoginAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_openaccount_login_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBaichuanOpenaccountLoginAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Name = ""
}

var poolTaobaoBaichuanOpenaccountLoginAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBaichuanOpenaccountLoginAPIResponse)
	},
}

// GetTaobaoBaichuanOpenaccountLoginAPIResponse 从 sync.Pool 获取 TaobaoBaichuanOpenaccountLoginAPIResponse
func GetTaobaoBaichuanOpenaccountLoginAPIResponse() *TaobaoBaichuanOpenaccountLoginAPIResponse {
	return poolTaobaoBaichuanOpenaccountLoginAPIResponse.Get().(*TaobaoBaichuanOpenaccountLoginAPIResponse)
}

// ReleaseTaobaoBaichuanOpenaccountLoginAPIResponse 将 TaobaoBaichuanOpenaccountLoginAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBaichuanOpenaccountLoginAPIResponse(v *TaobaoBaichuanOpenaccountLoginAPIResponse) {
	v.Reset()
	poolTaobaoBaichuanOpenaccountLoginAPIResponse.Put(v)
}
