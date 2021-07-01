package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
查询外部成本中心 
alitrip.btip.cost.center.query

查询外部成本中心
*/
func AlitripBtipCostCenterQuery(clt *core.SDKClient, req *btrip.AlitripBtipCostCenterQueryAPIRequest, session string) (*btrip.AlitripBtipCostCenterQueryAPIResponse, error) {
    var resp btrip.AlitripBtipCostCenterQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
