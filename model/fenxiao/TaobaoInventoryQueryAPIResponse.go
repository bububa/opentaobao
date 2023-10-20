package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoInventoryQueryAPIResponse 查询商品库存信息 API返回值
// taobao.inventory.query
//
// 建议使用新接口：tmall.inventory.query.forstore ，新ISV不推荐使用。
// 商家查询商品总体库存信息
type TaobaoInventoryQueryAPIResponse struct {
	model.CommonResponse
	TaobaoInventoryQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoInventoryQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoInventoryQueryAPIResponseModel).Reset()
}

// TaobaoInventoryQueryAPIResponseModel is 查询商品库存信息 成功返回结果
type TaobaoInventoryQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"inventory_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品总体库存信息
	ItemInventorys []InventorySum `json:"item_inventorys,omitempty" xml:"item_inventorys>inventory_sum,omitempty"`
	// 提示信息，提示不存在的后端商品
	TipInfos []TipInfo `json:"tip_infos,omitempty" xml:"tip_infos>tip_info,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoInventoryQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ItemInventorys = m.ItemInventorys[:0]
	m.TipInfos = m.TipInfos[:0]
}

var poolTaobaoInventoryQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoInventoryQueryAPIResponse)
	},
}

// GetTaobaoInventoryQueryAPIResponse 从 sync.Pool 获取 TaobaoInventoryQueryAPIResponse
func GetTaobaoInventoryQueryAPIResponse() *TaobaoInventoryQueryAPIResponse {
	return poolTaobaoInventoryQueryAPIResponse.Get().(*TaobaoInventoryQueryAPIResponse)
}

// ReleaseTaobaoInventoryQueryAPIResponse 将 TaobaoInventoryQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoInventoryQueryAPIResponse(v *TaobaoInventoryQueryAPIResponse) {
	v.Reset()
	poolTaobaoInventoryQueryAPIResponse.Put(v)
}
