package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
汽车票行业搜索接口 
alitrip.btrip.supplychain.bus.industry.search

汽车票行业搜索接口
*/
func AlitripBtripSupplychainBusIndustrySearch(clt *core.SDKClient, req *btrip.AlitripBtripSupplychainBusIndustrySearchAPIRequest, session string) (*btrip.AlitripBtripSupplychainBusIndustrySearchAPIResponse, error) {
    var resp btrip.AlitripBtripSupplychainBusIndustrySearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
