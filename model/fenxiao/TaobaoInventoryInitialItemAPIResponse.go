package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoInventoryInitialItemAPIResponse 商品库存初始化 API返回值
// taobao.inventory.initial.item
//
// 建议使用新接口：taobao.inventory.merchant.adjust ，该接口会逐步停用。
// 商家仓商品初始化在各个仓中库存
type TaobaoInventoryInitialItemAPIResponse struct {
	model.CommonResponse
	TaobaoInventoryInitialItemAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoInventoryInitialItemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoInventoryInitialItemAPIResponseModel).Reset()
}

// TaobaoInventoryInitialItemAPIResponseModel is 商品库存初始化 成功返回结果
type TaobaoInventoryInitialItemAPIResponseModel struct {
	XMLName xml.Name `xml:"inventory_initial_item_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 提示信息
	TipInfos []TipInfo `json:"tip_infos,omitempty" xml:"tip_infos>tip_info,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoInventoryInitialItemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TipInfos = m.TipInfos[:0]
}

var poolTaobaoInventoryInitialItemAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoInventoryInitialItemAPIResponse)
	},
}

// GetTaobaoInventoryInitialItemAPIResponse 从 sync.Pool 获取 TaobaoInventoryInitialItemAPIResponse
func GetTaobaoInventoryInitialItemAPIResponse() *TaobaoInventoryInitialItemAPIResponse {
	return poolTaobaoInventoryInitialItemAPIResponse.Get().(*TaobaoInventoryInitialItemAPIResponse)
}

// ReleaseTaobaoInventoryInitialItemAPIResponse 将 TaobaoInventoryInitialItemAPIResponse 保存到 sync.Pool
func ReleaseTaobaoInventoryInitialItemAPIResponse(v *TaobaoInventoryInitialItemAPIResponse) {
	v.Reset()
	poolTaobaoInventoryInitialItemAPIResponse.Put(v)
}
