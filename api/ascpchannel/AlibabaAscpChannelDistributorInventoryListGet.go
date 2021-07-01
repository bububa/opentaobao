package ascpchannel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ascpchannel"
)

/* 
批量查询渠道库存 
alibaba.ascp.channel.distributor.inventory.list.get

淘外分销批量查询渠道产品的库存
*/
func AlibabaAscpChannelDistributorInventoryListGet(clt *core.SDKClient, req *ascpchannel.AlibabaAscpChannelDistributorInventoryListGetAPIRequest, session string) (*ascpchannel.AlibabaAscpChannelDistributorInventoryListGetAPIResponse, error) {
    var resp ascpchannel.AlibabaAscpChannelDistributorInventoryListGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
