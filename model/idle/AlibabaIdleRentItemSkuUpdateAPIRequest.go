package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleRentItemSkuUpdateAPIRequest
更新/增加sku信息 API请求
alibaba.idle.rent.item.sku.update

更新/增加sku信息 */
type AlibabaIdleRentItemSkuUpdateAPIRequest struct {
	model.Params
	// sku信息，更新后skuId保持不变
	_sku *ItemSkuDto
}

// New
