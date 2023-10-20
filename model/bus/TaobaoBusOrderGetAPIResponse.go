package bus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusOrderGetAPIResponse 汽车票订单查询 API返回值
// taobao.bus.order.get
//
// 商家汽车票订单查询
type TaobaoBusOrderGetAPIResponse struct {
	model.CommonResponse
	TaobaoBusOrderGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBusOrderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBusOrderGetAPIResponseModel).Reset()
}

// TaobaoBusOrderGetAPIResponseModel is 汽车票订单查询 成功返回结果
type TaobaoBusOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单查询返回对象
	Result *B2BOrderQueryRp `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBusOrderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoBusOrderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBusOrderGetAPIResponse)
	},
}

// GetTaobaoBusOrderGetAPIResponse 从 sync.Pool 获取 TaobaoBusOrderGetAPIResponse
func GetTaobaoBusOrderGetAPIResponse() *TaobaoBusOrderGetAPIResponse {
	return poolTaobaoBusOrderGetAPIResponse.Get().(*TaobaoBusOrderGetAPIResponse)
}

// ReleaseTaobaoBusOrderGetAPIResponse 将 TaobaoBusOrderGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBusOrderGetAPIResponse(v *TaobaoBusOrderGetAPIResponse) {
	v.Reset()
	poolTaobaoBusOrderGetAPIResponse.Put(v)
}
