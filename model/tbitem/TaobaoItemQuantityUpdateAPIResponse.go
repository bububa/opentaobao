package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemQuantityUpdateAPIResponse 宝贝/SKU库存修改 API返回值
// taobao.item.quantity.update
//
// 提供按照全量或增量形式修改宝贝/SKU库存的功能
type TaobaoItemQuantityUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoItemQuantityUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoItemQuantityUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoItemQuantityUpdateAPIResponseModel).Reset()
}

// TaobaoItemQuantityUpdateAPIResponseModel is 宝贝/SKU库存修改 成功返回结果
type TaobaoItemQuantityUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"item_quantity_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// iid、numIid、num和modified，skus中每个sku的skuId、quantity和modified
	Item *Item `json:"item,omitempty" xml:"item,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoItemQuantityUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Item = nil
}

var poolTaobaoItemQuantityUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoItemQuantityUpdateAPIResponse)
	},
}

// GetTaobaoItemQuantityUpdateAPIResponse 从 sync.Pool 获取 TaobaoItemQuantityUpdateAPIResponse
func GetTaobaoItemQuantityUpdateAPIResponse() *TaobaoItemQuantityUpdateAPIResponse {
	return poolTaobaoItemQuantityUpdateAPIResponse.Get().(*TaobaoItemQuantityUpdateAPIResponse)
}

// ReleaseTaobaoItemQuantityUpdateAPIResponse 将 TaobaoItemQuantityUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoItemQuantityUpdateAPIResponse(v *TaobaoItemQuantityUpdateAPIResponse) {
	v.Reset()
	poolTaobaoItemQuantityUpdateAPIResponse.Put(v)
}
