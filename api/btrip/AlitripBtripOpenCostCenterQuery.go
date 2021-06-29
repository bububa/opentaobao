package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
查询成本中心 
alitrip.btrip.open.cost.center.query

查询成本中心
*/
func AlitripBtripOpenCostCenterQuery(clt *core.SDKClient, req *btrip.AlitripBtripOpenCostCenterQueryRequest, session string) (*btrip.AlitripBtripOpenCostCenterQueryAPIResponse, error) {
    var resp btrip.AlitripBtripOpenCostCenterQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
