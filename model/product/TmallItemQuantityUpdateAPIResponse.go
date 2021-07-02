package product

import (
	"encoding/xml"

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

// TmallItemQuantityUpdateAPIResponseModel is 天猫商品/SKU库存更新接口 成功返回结果
type TmallItemQuantityUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_quantity_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 库存更新结果，商品id
	QuantityUpdateResult string `json:"quantity_update_result,omitempty" xml:"quantity_update_result,omitempty"`
}
