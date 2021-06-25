package car

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/car"
)

/* 
CRS接送机订单列表搜索 
alitrip.travel.crsorder.search

提供给CRS商家搜索订单列表，仅返回订单号列表
*/
func AlitripTravelCrsorderSearch(clt *core.SDKClient, req *car.AlitripTravelCrsorderSearchRequest, session string) (*car.AlitripTravelCrsorderSearchResponse, error) {
    var resp car.AlitripTravelCrsorderSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
