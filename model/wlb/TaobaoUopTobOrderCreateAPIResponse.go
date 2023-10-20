package wlb

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUopTobOrderCreateAPIResponse ToB仓储发货 API返回值
// taobao.uop.tob.order.create
//
// ToB仓储发货
type TaobaoUopTobOrderCreateAPIResponse struct {
	model.CommonResponse
	TaobaoUopTobOrderCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUopTobOrderCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUopTobOrderCreateAPIResponseModel).Reset()
}

// TaobaoUopTobOrderCreateAPIResponseModel is ToB仓储发货 成功返回结果
type TaobaoUopTobOrderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"uop_tob_order_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单
	DeliveryOrders []DeliveryOrder `json:"delivery_orders,omitempty" xml:"delivery_orders>delivery_order,omitempty"`
	// flag
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUopTobOrderCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DeliveryOrders = m.DeliveryOrders[:0]
	m.Flag = ""
	m.Message = ""
}

var poolTaobaoUopTobOrderCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUopTobOrderCreateAPIResponse)
	},
}

// GetTaobaoUopTobOrderCreateAPIResponse 从 sync.Pool 获取 TaobaoUopTobOrderCreateAPIResponse
func GetTaobaoUopTobOrderCreateAPIResponse() *TaobaoUopTobOrderCreateAPIResponse {
	return poolTaobaoUopTobOrderCreateAPIResponse.Get().(*TaobaoUopTobOrderCreateAPIResponse)
}

// ReleaseTaobaoUopTobOrderCreateAPIResponse 将 TaobaoUopTobOrderCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUopTobOrderCreateAPIResponse(v *TaobaoUopTobOrderCreateAPIResponse) {
	v.Reset()
	poolTaobaoUopTobOrderCreateAPIResponse.Put(v)
}
