package bus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusLastplaceGetAPIResponse 获取目的地数据 API返回值
// taobao.bus.lastplace.get
//
// 传入城市 获取对应的目的地
type TaobaoBusLastplaceGetAPIResponse struct {
	model.CommonResponse
	TaobaoBusLastplaceGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBusLastplaceGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBusLastplaceGetAPIResponseModel).Reset()
}

// TaobaoBusLastplaceGetAPIResponseModel is 获取目的地数据 成功返回结果
type TaobaoBusLastplaceGetAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_lastplace_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 目的地返回结果
	Result *TaobaoBusLastplaceGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBusLastplaceGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoBusLastplaceGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBusLastplaceGetAPIResponse)
	},
}

// GetTaobaoBusLastplaceGetAPIResponse 从 sync.Pool 获取 TaobaoBusLastplaceGetAPIResponse
func GetTaobaoBusLastplaceGetAPIResponse() *TaobaoBusLastplaceGetAPIResponse {
	return poolTaobaoBusLastplaceGetAPIResponse.Get().(*TaobaoBusLastplaceGetAPIResponse)
}

// ReleaseTaobaoBusLastplaceGetAPIResponse 将 TaobaoBusLastplaceGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBusLastplaceGetAPIResponse(v *TaobaoBusLastplaceGetAPIResponse) {
	v.Reset()
	poolTaobaoBusLastplaceGetAPIResponse.Put(v)
}
