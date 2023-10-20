package car

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripCarOrderStatusAPIResponse 商家订单状态改变通知接口（神州专车接口） API返回值
// taobao.alitrip.car.order.status
//
// 商家订单状态改变通知接口，神州专车专用接口！
type TaobaoAlitripCarOrderStatusAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripCarOrderStatusAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripCarOrderStatusAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripCarOrderStatusAPIResponseModel).Reset()
}

// TaobaoAlitripCarOrderStatusAPIResponseModel is 商家订单状态改变通知接口（神州专车接口） 成功返回结果
type TaobaoAlitripCarOrderStatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_car_order_status_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 根据站点名称查询产品
	Result *TaobaoAlitripCarOrderStatusApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripCarOrderStatusAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripCarOrderStatusAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripCarOrderStatusAPIResponse)
	},
}

// GetTaobaoAlitripCarOrderStatusAPIResponse 从 sync.Pool 获取 TaobaoAlitripCarOrderStatusAPIResponse
func GetTaobaoAlitripCarOrderStatusAPIResponse() *TaobaoAlitripCarOrderStatusAPIResponse {
	return poolTaobaoAlitripCarOrderStatusAPIResponse.Get().(*TaobaoAlitripCarOrderStatusAPIResponse)
}

// ReleaseTaobaoAlitripCarOrderStatusAPIResponse 将 TaobaoAlitripCarOrderStatusAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripCarOrderStatusAPIResponse(v *TaobaoAlitripCarOrderStatusAPIResponse) {
	v.Reset()
	poolTaobaoAlitripCarOrderStatusAPIResponse.Put(v)
}
