package car

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripCarOrderQueryAPIResponse 飞猪订单状态查询接口 API返回值
// taobao.alitrip.car.order.query
//
// 提供给直连商家查询在飞猪平台上产生的订单
type TaobaoAlitripCarOrderQueryAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripCarOrderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripCarOrderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripCarOrderQueryAPIResponseModel).Reset()
}

// TaobaoAlitripCarOrderQueryAPIResponseModel is 飞猪订单状态查询接口 成功返回结果
type TaobaoAlitripCarOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_car_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单结果
	FirstResult *OrderQueryRsp `json:"first_result,omitempty" xml:"first_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripCarOrderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FirstResult = nil
}

var poolTaobaoAlitripCarOrderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripCarOrderQueryAPIResponse)
	},
}

// GetTaobaoAlitripCarOrderQueryAPIResponse 从 sync.Pool 获取 TaobaoAlitripCarOrderQueryAPIResponse
func GetTaobaoAlitripCarOrderQueryAPIResponse() *TaobaoAlitripCarOrderQueryAPIResponse {
	return poolTaobaoAlitripCarOrderQueryAPIResponse.Get().(*TaobaoAlitripCarOrderQueryAPIResponse)
}

// ReleaseTaobaoAlitripCarOrderQueryAPIResponse 将 TaobaoAlitripCarOrderQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripCarOrderQueryAPIResponse(v *TaobaoAlitripCarOrderQueryAPIResponse) {
	v.Reset()
	poolTaobaoAlitripCarOrderQueryAPIResponse.Put(v)
}
