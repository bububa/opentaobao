package pur

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaPurProductSyncAPIRequest
同步产品 API请求
alibaba.pur.product.sync

同步产品 */
type AlibabaPurProductSyncAPIRequest struct {
	model.Params
	// 产品对象
	_accessProductDtos []AccessProductDto
}

// New
