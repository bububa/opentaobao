package icbuproduct

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductInventoryUpdateAPIResponse icbu商品库存更新 API返回值
// alibaba.icbu.product.inventory.update
//
// 更新库存信息
type AlibabaIcbuProductInventoryUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuProductInventoryUpdateAPIResponseModel
}

// AlibabaIcbuProductInventoryUpdateAPIResponseModel is icbu商品库存更新 成功返回结果
type AlibabaIcbuProductInventoryUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_product_inventory_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Top返回对象
	Result *TopResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
