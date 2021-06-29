package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
三方市内用车申请单查询 
alitrip.btrip.city.car.apply.query

三方市内用车申请单查询
*/
func AlitripBtripCityCarApplyQuery(clt *core.SDKClient, req *btrip.AlitripBtripCityCarApplyQueryRequest, session string) (*btrip.AlitripBtripCityCarApplyQueryAPIResponse, error) {
    var resp btrip.AlitripBtripCityCarApplyQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
