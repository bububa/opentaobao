package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
链渠道中心淘外库存查询 API请求
alibaba.ascp.channel.distributor.inventory.get

此api为淘外分销的渠道产品库存查询标准api，淘外分销商专用
*/
type AlibabaAscpChannelDistributorInventoryGetRequest struct {
    model.Params
    // 入参
    _inventoryRequest   *ChannelInventoryQuery
}

// 初始化AlibabaAscpChannelDistributorInventoryGetRequest对象
func NewAlibabaAscpChannelDistributorInventoryGetRequest() *AlibabaAscpChannelDistributorInventoryGetRequest{
    return &AlibabaAscpChannelDistributorInventoryGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelDistributorInventoryGetRequest) GetApiMethodName() string {
    return "alibaba.ascp.channel.distributor.inventory.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelDistributorInventoryGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InventoryRequest Setter
// 入参
func (r *AlibabaAscpChannelDistributorInventoryGetRequest) SetInventoryRequest(_inventoryRequest *ChannelInventoryQuery) error {
    r._inventoryRequest = _inventoryRequest
    r.Set("inventory_request", _inventoryRequest)
    return nil
}

// InventoryRequest Getter
func (r AlibabaAscpChannelDistributorInventoryGetRequest) GetInventoryRequest() *ChannelInventoryQuery {
    return r._inventoryRequest
}
