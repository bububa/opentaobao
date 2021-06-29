package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
三方市内用车申请单审批 
alitrip.btrip.city.car.apply.approve

三方市内用车申请单审批
*/
func AlitripBtripCityCarApplyApprove(clt *core.SDKClient, req *btrip.AlitripBtripCityCarApplyApproveRequest, session string) (*btrip.AlitripBtripCityCarApplyApproveAPIResponse, error) {
    var resp btrip.AlitripBtripCityCarApplyApproveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
