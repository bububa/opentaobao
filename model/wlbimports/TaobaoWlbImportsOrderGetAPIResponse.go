package wlbimports

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbImportsOrderGetAPIResponse 物流订单获取 API返回值
// taobao.wlb.imports.order.get
//
// 一般进口物流订单获取
type TaobaoWlbImportsOrderGetAPIResponse struct {
	model.CommonResponse
	TaobaoWlbImportsOrderGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbImportsOrderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbImportsOrderGetAPIResponseModel).Reset()
}

// TaobaoWlbImportsOrderGetAPIResponseModel is 物流订单获取 成功返回结果
type TaobaoWlbImportsOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_imports_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 物流订单信息
	Orders []LocOrder `json:"orders,omitempty" xml:"orders>loc_order,omitempty"`
	// 搜索到的总数量
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbImportsOrderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Orders = m.Orders[:0]
	m.TotalResults = 0
}

var poolTaobaoWlbImportsOrderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbImportsOrderGetAPIResponse)
	},
}

// GetTaobaoWlbImportsOrderGetAPIResponse 从 sync.Pool 获取 TaobaoWlbImportsOrderGetAPIResponse
func GetTaobaoWlbImportsOrderGetAPIResponse() *TaobaoWlbImportsOrderGetAPIResponse {
	return poolTaobaoWlbImportsOrderGetAPIResponse.Get().(*TaobaoWlbImportsOrderGetAPIResponse)
}

// ReleaseTaobaoWlbImportsOrderGetAPIResponse 将 TaobaoWlbImportsOrderGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbImportsOrderGetAPIResponse(v *TaobaoWlbImportsOrderGetAPIResponse) {
	v.Reset()
	poolTaobaoWlbImportsOrderGetAPIResponse.Put(v)
}
