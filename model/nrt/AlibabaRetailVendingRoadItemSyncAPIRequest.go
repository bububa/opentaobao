package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailVendingRoadItemSyncAPIRequest 贩卖机库存商品同步 API请求
// alibaba.retail.vending.road.item.sync
//
// 贩卖机库存商品同步
type AlibabaRetailVendingRoadItemSyncAPIRequest struct {
	model.Params
	// 入参
	_roadItemSync *RoadItemSyncDto
}

// NewAlibabaRetailVendingRoadItemSyncRequest 初始化AlibabaRetailVendingRoadItemSyncAPIRequest对象
func NewAlibabaRetailVendingRoadItemSyncRequest() *AlibabaRetailVendingRoadItemSyncAPIRequest {
	return &AlibabaRetailVendingRoadItemSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailVendingRoadItemSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.vending.road.item.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailVendingRoadItemSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRoadItemSync is RoadItemSync Setter
// 入参
func (r *AlibabaRetailVendingRoadItemSyncAPIRequest) SetRoadItemSync(_roadItemSync *RoadItemSyncDto) error {
	r._roadItemSync = _roadItemSync
	r.Set("road_item_sync", _roadItemSync)
	return nil
}

// GetRoadItemSync RoadItemSync Getter
func (r AlibabaRetailVendingRoadItemSyncAPIRequest) GetRoadItemSync() *RoadItemSyncDto {
	return r._roadItemSync
}
