package icbuproduct

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuProductInventoryUpdateAPIRequest
icbu商品库存更新 API请求
alibaba.icbu.product.inventory.update

更新库存信息 */
type AlibabaIcbuProductInventoryUpdateAPIRequest struct {
	model.Params
	// 更新请求
	_requestParam *ProductInventoryRequest
}

// New
