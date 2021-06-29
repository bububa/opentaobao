package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
月账单数据查询 
alitrip.btrip.monthbill.url.get

月账单数据查询
*/
func AlitripBtripMonthbillUrlGet(clt *core.SDKClient, req *btrip.AlitripBtripMonthbillUrlGetRequest, session string) (*btrip.AlitripBtripMonthbillUrlGetAPIResponse, error) {
    var resp btrip.AlitripBtripMonthbillUrlGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
