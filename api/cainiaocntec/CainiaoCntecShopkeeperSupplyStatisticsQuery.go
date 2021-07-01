package cainiaocntec

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/cainiaocntec"
)

/* 
团购业务供货商查询门店统计数据 
cainiao.cntec.shopkeeper.supply.statistics.query

查询门店售卖商品统计数据
*/
func CainiaoCntecShopkeeperSupplyStatisticsQuery(clt *core.SDKClient, req *cainiaocntec.CainiaoCntecShopkeeperSupplyStatisticsQueryAPIRequest, session string) (*cainiaocntec.CainiaoCntecShopkeeperSupplyStatisticsQueryAPIResponse, error) {
    var resp cainiaocntec.CainiaoCntecShopkeeperSupplyStatisticsQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
