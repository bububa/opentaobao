package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
火车票行业搜索接口 
alitrip.btrip.supplychain.train.industry.search

【商旅】火车票行业搜索接口
*/
func AlitripBtripSupplychainTrainIndustrySearch(clt *core.SDKClient, req *btrip.AlitripBtripSupplychainTrainIndustrySearchAPIRequest, session string) (*btrip.AlitripBtripSupplychainTrainIndustrySearchAPIResponse, error) {
    var resp btrip.AlitripBtripSupplychainTrainIndustrySearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
