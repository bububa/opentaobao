package baichuan

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanOpenaccountRegistercodeCheckAPIResponse 百川检查注册验证码 API返回值
// taobao.baichuan.openaccount.registercode.check
//
// 百川检查注册验证码
type TaobaoBaichuanOpenaccountRegistercodeCheckAPIResponse struct {
	model.CommonResponse
	TaobaoBaichuanOpenaccountRegistercodeCheckAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBaichuanOpenaccountRegistercodeCheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBaichuanOpenaccountRegistercodeCheckAPIResponseModel).Reset()
}

// TaobaoBaichuanOpenaccountRegistercodeCheckAPIResponseModel is 百川检查注册验证码 成功返回结果
type TaobaoBaichuanOpenaccountRegistercodeCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_openaccount_registercode_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBaichuanOpenaccountRegistercodeCheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Name = ""
}

var poolTaobaoBaichuanOpenaccountRegistercodeCheckAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBaichuanOpenaccountRegistercodeCheckAPIResponse)
	},
}

// GetTaobaoBaichuanOpenaccountRegistercodeCheckAPIResponse 从 sync.Pool 获取 TaobaoBaichuanOpenaccountRegistercodeCheckAPIResponse
func GetTaobaoBaichuanOpenaccountRegistercodeCheckAPIResponse() *TaobaoBaichuanOpenaccountRegistercodeCheckAPIResponse {
	return poolTaobaoBaichuanOpenaccountRegistercodeCheckAPIResponse.Get().(*TaobaoBaichuanOpenaccountRegistercodeCheckAPIResponse)
}

// ReleaseTaobaoBaichuanOpenaccountRegistercodeCheckAPIResponse 将 TaobaoBaichuanOpenaccountRegistercodeCheckAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBaichuanOpenaccountRegistercodeCheckAPIResponse(v *TaobaoBaichuanOpenaccountRegistercodeCheckAPIResponse) {
	v.Reset()
	poolTaobaoBaichuanOpenaccountRegistercodeCheckAPIResponse.Put(v)
}
