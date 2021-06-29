package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询渠道库存 API请求
alibaba.ascp.channel.distributor.inventory.list.get

淘外分销批量查询渠道产品的库存
*/
type AlibabaAscpChannelDistributorInventoryListGetRequest struct {
    model.Params
    // 系统自动生成
    _inventoryRequest   *BatchChannelInventoryQuery
}

// 初始化AlibabaAscpChannelDistributorInventoryListGetRequest对象
func NewAlibabaAscpChannelDistributorInventoryListGetRequest() *AlibabaAscpChannelDistributorInventoryListGetRequest{
    return &AlibabaAscpChannelDistributorInventoryListGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelDistributorInventoryListGetRequest) GetApiMethodName() string {
    return "alibaba.ascp.channel.distributor.inventory.list.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelDistributorInventoryListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InventoryRequest Setter
// 系统自动生成
func (r *AlibabaAscpChannelDistributorInventoryListGetRequest) SetInventoryRequest(_inventoryRequest *BatchChannelInventoryQuery) error {
    r._inventoryRequest = _inventoryRequest
    r.Set("inventory_request", _inventoryRequest)
    return nil
}

// InventoryRequest Getter
func (r AlibabaAscpChannelDistributorInventoryListGetRequest) GetInventoryRequest() *BatchChannelInventoryQuery {
    return r._inventoryRequest
}
