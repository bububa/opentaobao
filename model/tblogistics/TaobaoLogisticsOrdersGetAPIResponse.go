package tblogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsOrdersGetAPIResponse 批量查询物流订单 API返回值
// taobao.logistics.orders.get
//
// 批量查询物流订单。
type TaobaoLogisticsOrdersGetAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsOrdersGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsOrdersGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsOrdersGetAPIResponseModel).Reset()
}

// TaobaoLogisticsOrdersGetAPIResponseModel is 批量查询物流订单 成功返回结果
type TaobaoLogisticsOrdersGetAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_orders_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 获取的物流订单详情列表 &lt;br/&gt;返回的Shipping包含的具体信息为入参fields请求的字段信息
	Shippings []Shipping `json:"shippings,omitempty" xml:"shippings>shipping,omitempty"`
	// 搜索到的物流订单列表总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsOrdersGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Shippings = m.Shippings[:0]
	m.TotalResults = 0
}

var poolTaobaoLogisticsOrdersGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsOrdersGetAPIResponse)
	},
}

// GetTaobaoLogisticsOrdersGetAPIResponse 从 sync.Pool 获取 TaobaoLogisticsOrdersGetAPIResponse
func GetTaobaoLogisticsOrdersGetAPIResponse() *TaobaoLogisticsOrdersGetAPIResponse {
	return poolTaobaoLogisticsOrdersGetAPIResponse.Get().(*TaobaoLogisticsOrdersGetAPIResponse)
}

// ReleaseTaobaoLogisticsOrdersGetAPIResponse 将 TaobaoLogisticsOrdersGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsOrdersGetAPIResponse(v *TaobaoLogisticsOrdersGetAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsOrdersGetAPIResponse.Put(v)
}
