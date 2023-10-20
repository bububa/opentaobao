package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoOrdersGetAPIResponse 查询采购单信息 API返回值
// taobao.fenxiao.orders.get
//
// 查询代销采购单单据。
//
// 1. 支持商家按照供应商、分销商两种角色来查询数据。如果没有指定角色角色，系统会自动判断，此时如果商家存在供应商、分销商两种角色时，按照供应商角色查询。
// 2. 同时此接口还可以查询除供销经销外的其他经营模式的数据。如果需要查询供销经销单据请参考接口：taobao.fenxiao.dealer.requisitionorder.query
//
// 3. 发货请调用物流API中的发货接口taobao.logistics.offline.send 进行发货，需要注意的是这里是供应商发货，因此调发货接口时需要传人供应商账号对应的sessionkey，tid 需传入供销平台的采购单（即fenxiao_id  分销流水号）)。
type TaobaoFenxiaoOrdersGetAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoOrdersGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFenxiaoOrdersGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFenxiaoOrdersGetAPIResponseModel).Reset()
}

// TaobaoFenxiaoOrdersGetAPIResponseModel is 查询采购单信息 成功返回结果
type TaobaoFenxiaoOrdersGetAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_orders_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 采购单及子采购单信息。 返回 PurchaseOrder 包含的字段信息。
	PurchaseOrders []TopDpOrderDo `json:"purchase_orders,omitempty" xml:"purchase_orders>top_dp_order_do,omitempty"`
	// 查询到的采购单记录总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFenxiaoOrdersGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.PurchaseOrders = m.PurchaseOrders[:0]
	m.TotalResults = 0
}

var poolTaobaoFenxiaoOrdersGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoOrdersGetAPIResponse)
	},
}

// GetTaobaoFenxiaoOrdersGetAPIResponse 从 sync.Pool 获取 TaobaoFenxiaoOrdersGetAPIResponse
func GetTaobaoFenxiaoOrdersGetAPIResponse() *TaobaoFenxiaoOrdersGetAPIResponse {
	return poolTaobaoFenxiaoOrdersGetAPIResponse.Get().(*TaobaoFenxiaoOrdersGetAPIResponse)
}

// ReleaseTaobaoFenxiaoOrdersGetAPIResponse 将 TaobaoFenxiaoOrdersGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFenxiaoOrdersGetAPIResponse(v *TaobaoFenxiaoOrdersGetAPIResponse) {
	v.Reset()
	poolTaobaoFenxiaoOrdersGetAPIResponse.Put(v)
}
