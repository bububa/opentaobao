package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemQuantityUpdateAPIResponse 天猫商品/SKU库存更新接口 API返回值
// tmall.item.quantity.update
//
// 天猫商品/SKU库存更新接口；支持商品库存更新；支持同一商品下的SKU批量更新。
type TmallItemQuantityUpdateAPIResponse struct {
	model.CommonResponse
	TmallItemQuantityUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallItemQuantityUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallItemQuantityUpdateAPIResponseModel).Reset()
}

// TmallItemQuantityUpdateAPIResponseModel is 天猫商品/SKU库存更新接口 成功返回结果
type TmallItemQuantityUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_quantity_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 库存更新结果，商品id
	QuantityUpdateResult string `json:"quantity_update_result,omitempty" xml:"quantity_update_result,omitempty"`
}

// Reset 清空结构体
func (m *TmallItemQuantityUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.QuantityUpdateResult = ""
}

var poolTmallItemQuantityUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallItemQuantityUpdateAPIResponse)
	},
}

// GetTmallItemQuantityUpdateAPIResponse 从 sync.Pool 获取 TmallItemQuantityUpdateAPIResponse
func GetTmallItemQuantityUpdateAPIResponse() *TmallItemQuantityUpdateAPIResponse {
	return poolTmallItemQuantityUpdateAPIResponse.Get().(*TmallItemQuantityUpdateAPIResponse)
}

// ReleaseTmallItemQuantityUpdateAPIResponse 将 TmallItemQuantityUpdateAPIResponse 保存到 sync.Pool
func ReleaseTmallItemQuantityUpdateAPIResponse(v *TmallItemQuantityUpdateAPIResponse) {
	v.Reset()
	poolTmallItemQuantityUpdateAPIResponse.Put(v)
}
