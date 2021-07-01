package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailVendingRoadItemSyncAPIRequest
贩卖机库存商品同步 API请求
alibaba.retail.vending.road.item.sync

贩卖机库存商品同步 */
type AlibabaRetailVendingRoadItemSyncAPIRequest struct {
	model.Params
	// 入参
	_roadItemSync *RoadItemSyncDto
}

// New
