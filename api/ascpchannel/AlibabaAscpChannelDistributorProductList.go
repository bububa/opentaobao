package ascpchannel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ascpchannel"
)

/* 
供应链渠道中心淘外分销品批量查询(分销商专用) 
alibaba.ascp.channel.distributor.product.list

此api为淘外分销的品批量查询标准api，淘外分销商专用
*/
func AlibabaAscpChannelDistributorProductList(clt *core.SDKClient, req *ascpchannel.AlibabaAscpChannelDistributorProductListAPIRequest, session string) (*ascpchannel.AlibabaAscpChannelDistributorProductListAPIResponse, error) {
    var resp ascpchannel.AlibabaAscpChannelDistributorProductListAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
