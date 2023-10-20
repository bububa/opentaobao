package baichuan

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanOpenaccountLogindoublecheckAPIResponse 百川登录二次验证 API返回值
// taobao.baichuan.openaccount.logindoublecheck
//
// 百川登录二次验证
type TaobaoBaichuanOpenaccountLogindoublecheckAPIResponse struct {
	model.CommonResponse
	TaobaoBaichuanOpenaccountLogindoublecheckAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBaichuanOpenaccountLogindoublecheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBaichuanOpenaccountLogindoublecheckAPIResponseModel).Reset()
}

// TaobaoBaichuanOpenaccountLogindoublecheckAPIResponseModel is 百川登录二次验证 成功返回结果
type TaobaoBaichuanOpenaccountLogindoublecheckAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_openaccount_logindoublecheck_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBaichuanOpenaccountLogindoublecheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Name = ""
}

var poolTaobaoBaichuanOpenaccountLogindoublecheckAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBaichuanOpenaccountLogindoublecheckAPIResponse)
	},
}

// GetTaobaoBaichuanOpenaccountLogindoublecheckAPIResponse 从 sync.Pool 获取 TaobaoBaichuanOpenaccountLogindoublecheckAPIResponse
func GetTaobaoBaichuanOpenaccountLogindoublecheckAPIResponse() *TaobaoBaichuanOpenaccountLogindoublecheckAPIResponse {
	return poolTaobaoBaichuanOpenaccountLogindoublecheckAPIResponse.Get().(*TaobaoBaichuanOpenaccountLogindoublecheckAPIResponse)
}

// ReleaseTaobaoBaichuanOpenaccountLogindoublecheckAPIResponse 将 TaobaoBaichuanOpenaccountLogindoublecheckAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBaichuanOpenaccountLogindoublecheckAPIResponse(v *TaobaoBaichuanOpenaccountLogindoublecheckAPIResponse) {
	v.Reset()
	poolTaobaoBaichuanOpenaccountLogindoublecheckAPIResponse.Put(v)
}
