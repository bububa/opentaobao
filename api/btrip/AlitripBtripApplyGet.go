package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
获取单个审批单 
alitrip.btrip.apply.get

获取单个审批单的详情数据
*/
func AlitripBtripApplyGet(clt *core.SDKClient, req *btrip.AlitripBtripApplyGetRequest, session string) (*btrip.AlitripBtripApplyGetAPIResponse, error) {
    var resp btrip.AlitripBtripApplyGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
