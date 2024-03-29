package wlb

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbItemBatchQueryAPIResponse 批次库存查询接口 API返回值
// taobao.wlb.item.batch.query
//
// 根据用户id，item id list和store code来查询商品库存信息和批次信息
type TaobaoWlbItemBatchQueryAPIResponse struct {
	model.CommonResponse
	TaobaoWlbItemBatchQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbItemBatchQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbItemBatchQueryAPIResponseModel).Reset()
}

// TaobaoWlbItemBatchQueryAPIResponseModel is 批次库存查询接口 成功返回结果
type TaobaoWlbItemBatchQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_item_batch_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品库存及批次信息查询结果
	ItemInventoryBatchList []WlbItemBatchInventory `json:"item_inventory_batch_list,omitempty" xml:"item_inventory_batch_list>wlb_item_batch_inventory,omitempty"`
	// 返回结果记录的数量
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbItemBatchQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ItemInventoryBatchList = m.ItemInventoryBatchList[:0]
	m.TotalCount = 0
}

var poolTaobaoWlbItemBatchQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbItemBatchQueryAPIResponse)
	},
}

// GetTaobaoWlbItemBatchQueryAPIResponse 从 sync.Pool 获取 TaobaoWlbItemBatchQueryAPIResponse
func GetTaobaoWlbItemBatchQueryAPIResponse() *TaobaoWlbItemBatchQueryAPIResponse {
	return poolTaobaoWlbItemBatchQueryAPIResponse.Get().(*TaobaoWlbItemBatchQueryAPIResponse)
}

// ReleaseTaobaoWlbItemBatchQueryAPIResponse 将 TaobaoWlbItemBatchQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbItemBatchQueryAPIResponse(v *TaobaoWlbItemBatchQueryAPIResponse) {
	v.Reset()
	poolTaobaoWlbItemBatchQueryAPIResponse.Put(v)
}
