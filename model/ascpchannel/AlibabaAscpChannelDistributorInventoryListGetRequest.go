package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询渠道库存 APIRequest
alibaba.ascp.channel.distributor.inventory.list.get

淘外分销批量查询渠道产品的库存
*/
type AlibabaAscpChannelDistributorInventoryListGetRequest struct {
    model.Params

    // 系统自动生成
    inventoryRequest   *BatchChannelInventoryQuery 

}

func NewAlibabaAscpChannelDistributorInventoryListGetRequest() *AlibabaAscpChannelDistributorInventoryListGetRequest{
    return &AlibabaAscpChannelDistributorInventoryListGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpChannelDistributorInventoryListGetRequest) GetApiMethodName() string {
    return "alibaba.ascp.channel.distributor.inventory.list.get"
}

func (r AlibabaAscpChannelDistributorInventoryListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpChannelDistributorInventoryListGetRequest) SetInventoryRequest(inventoryRequest *BatchChannelInventoryQuery) error {
    r.inventoryRequest = inventoryRequest
    r.Set("inventory_request", inventoryRequest)
    return nil
}

func (r AlibabaAscpChannelDistributorInventoryListGetRequest) GetInventoryRequest() *BatchChannelInventoryQuery {
    return r.inventoryRequest
}

