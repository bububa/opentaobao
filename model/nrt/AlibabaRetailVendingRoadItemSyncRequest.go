package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
贩卖机库存商品同步 APIRequest
alibaba.retail.vending.road.item.sync

贩卖机库存商品同步
*/
type AlibabaRetailVendingRoadItemSyncRequest struct {
    model.Params

    // 入参
    roadItemSync   *RoadItemSyncDto 

}

func NewAlibabaRetailVendingRoadItemSyncRequest() *AlibabaRetailVendingRoadItemSyncRequest{
    return &AlibabaRetailVendingRoadItemSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaRetailVendingRoadItemSyncRequest) GetApiMethodName() string {
    return "alibaba.retail.vending.road.item.sync"
}

func (r AlibabaRetailVendingRoadItemSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaRetailVendingRoadItemSyncRequest) SetRoadItemSync(roadItemSync *RoadItemSyncDto) error {
    r.roadItemSync = roadItemSync
    r.Set("road_item_sync", roadItemSync)
    return nil
}

func (r AlibabaRetailVendingRoadItemSyncRequest) GetRoadItemSync() *RoadItemSyncDto {
    return r.roadItemSync
}

