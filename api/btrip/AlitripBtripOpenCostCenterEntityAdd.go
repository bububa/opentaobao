package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
增加成本中心人员信息 
alitrip.btrip.open.cost.center.entity.add

增加成本中心人员信息
*/
func AlitripBtripOpenCostCenterEntityAdd(clt *core.SDKClient, req *btrip.AlitripBtripOpenCostCenterEntityAddRequest, session string) (*btrip.AlitripBtripOpenCostCenterEntityAddAPIResponse, error) {
    var resp btrip.AlitripBtripOpenCostCenterEntityAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}