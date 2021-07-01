package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
机票城市搜索 
alitrip.btrip.flight.city.suggest

提供机票城市搜索接口，提高OA用户对接效率
*/
func AlitripBtripFlightCitySuggest(clt *core.SDKClient, req *btrip.AlitripBtripFlightCitySuggestAPIRequest, session string) (*btrip.AlitripBtripFlightCitySuggestAPIResponse, error) {
    var resp btrip.AlitripBtripFlightCitySuggestAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
