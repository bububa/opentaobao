package baichuan

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanOpenaccountRegisterAPIResponse 百川账号注册 API返回值
// taobao.baichuan.openaccount.register
//
// 百川账号注册
type TaobaoBaichuanOpenaccountRegisterAPIResponse struct {
	model.CommonResponse
	TaobaoBaichuanOpenaccountRegisterAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBaichuanOpenaccountRegisterAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBaichuanOpenaccountRegisterAPIResponseModel).Reset()
}

// TaobaoBaichuanOpenaccountRegisterAPIResponseModel is 百川账号注册 成功返回结果
type TaobaoBaichuanOpenaccountRegisterAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_openaccount_register_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBaichuanOpenaccountRegisterAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Name = ""
}

var poolTaobaoBaichuanOpenaccountRegisterAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBaichuanOpenaccountRegisterAPIResponse)
	},
}

// GetTaobaoBaichuanOpenaccountRegisterAPIResponse 从 sync.Pool 获取 TaobaoBaichuanOpenaccountRegisterAPIResponse
func GetTaobaoBaichuanOpenaccountRegisterAPIResponse() *TaobaoBaichuanOpenaccountRegisterAPIResponse {
	return poolTaobaoBaichuanOpenaccountRegisterAPIResponse.Get().(*TaobaoBaichuanOpenaccountRegisterAPIResponse)
}

// ReleaseTaobaoBaichuanOpenaccountRegisterAPIResponse 将 TaobaoBaichuanOpenaccountRegisterAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBaichuanOpenaccountRegisterAPIResponse(v *TaobaoBaichuanOpenaccountRegisterAPIResponse) {
	v.Reset()
	poolTaobaoBaichuanOpenaccountRegisterAPIResponse.Put(v)
}
