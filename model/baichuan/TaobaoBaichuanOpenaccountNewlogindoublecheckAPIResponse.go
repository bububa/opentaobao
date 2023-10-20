package baichuan

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanOpenaccountNewlogindoublecheckAPIResponse 百川新登录二次验证 API返回值
// taobao.baichuan.openaccount.newlogindoublecheck
//
// 百川新登录二次验证
type TaobaoBaichuanOpenaccountNewlogindoublecheckAPIResponse struct {
	model.CommonResponse
	TaobaoBaichuanOpenaccountNewlogindoublecheckAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBaichuanOpenaccountNewlogindoublecheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBaichuanOpenaccountNewlogindoublecheckAPIResponseModel).Reset()
}

// TaobaoBaichuanOpenaccountNewlogindoublecheckAPIResponseModel is 百川新登录二次验证 成功返回结果
type TaobaoBaichuanOpenaccountNewlogindoublecheckAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_openaccount_newlogindoublecheck_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBaichuanOpenaccountNewlogindoublecheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Name = ""
}

var poolTaobaoBaichuanOpenaccountNewlogindoublecheckAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBaichuanOpenaccountNewlogindoublecheckAPIResponse)
	},
}

// GetTaobaoBaichuanOpenaccountNewlogindoublecheckAPIResponse 从 sync.Pool 获取 TaobaoBaichuanOpenaccountNewlogindoublecheckAPIResponse
func GetTaobaoBaichuanOpenaccountNewlogindoublecheckAPIResponse() *TaobaoBaichuanOpenaccountNewlogindoublecheckAPIResponse {
	return poolTaobaoBaichuanOpenaccountNewlogindoublecheckAPIResponse.Get().(*TaobaoBaichuanOpenaccountNewlogindoublecheckAPIResponse)
}

// ReleaseTaobaoBaichuanOpenaccountNewlogindoublecheckAPIResponse 将 TaobaoBaichuanOpenaccountNewlogindoublecheckAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBaichuanOpenaccountNewlogindoublecheckAPIResponse(v *TaobaoBaichuanOpenaccountNewlogindoublecheckAPIResponse) {
	v.Reset()
	poolTaobaoBaichuanOpenaccountNewlogindoublecheckAPIResponse.Put(v)
}
