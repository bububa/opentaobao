package wlb

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbInventorylogQueryAPIResponse 根据商品ID查询所有库存变更记录 API返回值
// taobao.wlb.inventorylog.query
//
// 通过商品ID等几个条件来分页查询库存变更记录
type TaobaoWlbInventorylogQueryAPIResponse struct {
	model.CommonResponse
	TaobaoWlbInventorylogQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbInventorylogQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbInventorylogQueryAPIResponseModel).Reset()
}

// TaobaoWlbInventorylogQueryAPIResponseModel is 根据商品ID查询所有库存变更记录 成功返回结果
type TaobaoWlbInventorylogQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_inventorylog_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 库存变更记录
	InventoryLogList []WlbItemInventoryLog `json:"inventory_log_list,omitempty" xml:"inventory_log_list>wlb_item_inventory_log,omitempty"`
	// 记录数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbInventorylogQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.InventoryLogList = m.InventoryLogList[:0]
	m.TotalCount = 0
}

var poolTaobaoWlbInventorylogQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbInventorylogQueryAPIResponse)
	},
}

// GetTaobaoWlbInventorylogQueryAPIResponse 从 sync.Pool 获取 TaobaoWlbInventorylogQueryAPIResponse
func GetTaobaoWlbInventorylogQueryAPIResponse() *TaobaoWlbInventorylogQueryAPIResponse {
	return poolTaobaoWlbInventorylogQueryAPIResponse.Get().(*TaobaoWlbInventorylogQueryAPIResponse)
}

// ReleaseTaobaoWlbInventorylogQueryAPIResponse 将 TaobaoWlbInventorylogQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbInventorylogQueryAPIResponse(v *TaobaoWlbInventorylogQueryAPIResponse) {
	v.Reset()
	poolTaobaoWlbInventorylogQueryAPIResponse.Put(v)
}
