package wlb

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbOrderitemPageGetAPIResponse 分页查询物流宝订单商品详情 API返回值
// taobao.wlb.orderitem.page.get
//
// 分页查询物流宝订单商品详情
type TaobaoWlbOrderitemPageGetAPIResponse struct {
	model.CommonResponse
	TaobaoWlbOrderitemPageGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbOrderitemPageGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbOrderitemPageGetAPIResponseModel).Reset()
}

// TaobaoWlbOrderitemPageGetAPIResponseModel is 分页查询物流宝订单商品详情 成功返回结果
type TaobaoWlbOrderitemPageGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_orderitem_page_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	OrderItemList []WlbOrderItem `json:"order_item_list,omitempty" xml:"order_item_list>wlb_order_item,omitempty"`
	// 总数量
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbOrderitemPageGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OrderItemList = m.OrderItemList[:0]
	m.TotalCount = 0
}

var poolTaobaoWlbOrderitemPageGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbOrderitemPageGetAPIResponse)
	},
}

// GetTaobaoWlbOrderitemPageGetAPIResponse 从 sync.Pool 获取 TaobaoWlbOrderitemPageGetAPIResponse
func GetTaobaoWlbOrderitemPageGetAPIResponse() *TaobaoWlbOrderitemPageGetAPIResponse {
	return poolTaobaoWlbOrderitemPageGetAPIResponse.Get().(*TaobaoWlbOrderitemPageGetAPIResponse)
}

// ReleaseTaobaoWlbOrderitemPageGetAPIResponse 将 TaobaoWlbOrderitemPageGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbOrderitemPageGetAPIResponse(v *TaobaoWlbOrderitemPageGetAPIResponse) {
	v.Reset()
	poolTaobaoWlbOrderitemPageGetAPIResponse.Put(v)
}
