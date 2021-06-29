package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
贩卖机库存商品同步 API请求
alibaba.retail.vending.road.item.sync

贩卖机库存商品同步
*/
type AlibabaRetailVendingRoadItemSyncRequest struct {
    model.Params
    // 入参
    roadItemSync   *RoadItemSyncDto
}

// 初始化AlibabaRetailVendingRoadItemSyncRequest对象
func NewAlibabaRetailVendingRoadItemSyncRequest() *AlibabaRetailVendingRoadItemSyncRequest{
    return &AlibabaRetailVendingRoadItemSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailVendingRoadItemSyncRequest) GetApiMethodName() string {
    return "alibaba.retail.vending.road.item.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailVendingRoadItemSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RoadItemSync Setter
// 入参
func (r *AlibabaRetailVendingRoadItemSyncRequest) SetRoadItemSync(roadItemSync *RoadItemSyncDto) error {
    r.roadItemSync = roadItemSync
    r.Set("road_item_sync", roadItemSync)
    return nil
}

// RoadItemSync Getter
func (r AlibabaRetailVendingRoadItemSyncRequest) GetRoadItemSync() *RoadItemSyncDto {
    return r.roadItemSync
}
