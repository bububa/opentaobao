package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
删除外部成本中心人员信息 
alitrip.btrip.cost.center.entity.delete

删除外部成本中心人员信息
*/
func AlitripBtripCostCenterEntityDelete(clt *core.SDKClient, req *btrip.AlitripBtripCostCenterEntityDeleteRequest, session string) (*btrip.AlitripBtripCostCenterEntityDeleteAPIResponse, error) {
    var resp btrip.AlitripBtripCostCenterEntityDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
