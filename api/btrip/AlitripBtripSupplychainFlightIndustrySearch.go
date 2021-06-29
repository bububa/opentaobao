package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
机票行业搜索接口 
alitrip.btrip.supplychain.flight.industry.search

【商旅】机票行业搜索
*/
func AlitripBtripSupplychainFlightIndustrySearch(clt *core.SDKClient, req *btrip.AlitripBtripSupplychainFlightIndustrySearchRequest, session string) (*btrip.AlitripBtripSupplychainFlightIndustrySearchAPIResponse, error) {
    var resp btrip.AlitripBtripSupplychainFlightIndustrySearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
