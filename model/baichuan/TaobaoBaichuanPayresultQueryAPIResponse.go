package baichuan

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanPayresultQueryAPIResponse 百川支付完成回调 API返回值
// taobao.baichuan.payresult.query
//
// 百川支付完成回调
type TaobaoBaichuanPayresultQueryAPIResponse struct {
	model.CommonResponse
	TaobaoBaichuanPayresultQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBaichuanPayresultQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBaichuanPayresultQueryAPIResponseModel).Reset()
}

// TaobaoBaichuanPayresultQueryAPIResponseModel is 百川支付完成回调 成功返回结果
type TaobaoBaichuanPayresultQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_payresult_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBaichuanPayresultQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Name = ""
}

var poolTaobaoBaichuanPayresultQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBaichuanPayresultQueryAPIResponse)
	},
}

// GetTaobaoBaichuanPayresultQueryAPIResponse 从 sync.Pool 获取 TaobaoBaichuanPayresultQueryAPIResponse
func GetTaobaoBaichuanPayresultQueryAPIResponse() *TaobaoBaichuanPayresultQueryAPIResponse {
	return poolTaobaoBaichuanPayresultQueryAPIResponse.Get().(*TaobaoBaichuanPayresultQueryAPIResponse)
}

// ReleaseTaobaoBaichuanPayresultQueryAPIResponse 将 TaobaoBaichuanPayresultQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBaichuanPayresultQueryAPIResponse(v *TaobaoBaichuanPayresultQueryAPIResponse) {
	v.Reset()
	poolTaobaoBaichuanPayresultQueryAPIResponse.Put(v)
}
