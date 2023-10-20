package wlb

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbOrderPageGetAPIResponse 分页查询物流宝订单 API返回值
// taobao.wlb.order.page.get
//
// 分页查询物流宝订单
type TaobaoWlbOrderPageGetAPIResponse struct {
	model.CommonResponse
	TaobaoWlbOrderPageGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbOrderPageGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbOrderPageGetAPIResponseModel).Reset()
}

// TaobaoWlbOrderPageGetAPIResponseModel is 分页查询物流宝订单 成功返回结果
type TaobaoWlbOrderPageGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_order_page_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 物流宝订单对象
	OrderList []WlbOrder `json:"order_list,omitempty" xml:"order_list>wlb_order,omitempty"`
	// 总条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbOrderPageGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OrderList = m.OrderList[:0]
	m.TotalCount = 0
}

var poolTaobaoWlbOrderPageGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbOrderPageGetAPIResponse)
	},
}

// GetTaobaoWlbOrderPageGetAPIResponse 从 sync.Pool 获取 TaobaoWlbOrderPageGetAPIResponse
func GetTaobaoWlbOrderPageGetAPIResponse() *TaobaoWlbOrderPageGetAPIResponse {
	return poolTaobaoWlbOrderPageGetAPIResponse.Get().(*TaobaoWlbOrderPageGetAPIResponse)
}

// ReleaseTaobaoWlbOrderPageGetAPIResponse 将 TaobaoWlbOrderPageGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbOrderPageGetAPIResponse(v *TaobaoWlbOrderPageGetAPIResponse) {
	v.Reset()
	poolTaobaoWlbOrderPageGetAPIResponse.Put(v)
}
