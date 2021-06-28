package mos

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mos"
)

/* 
云闪付、银行卡销售数据回传接口 
alibaba.mos.oc.trade.syncbanksale

云闪付、银行卡销售数据回传
*/
func AlibabaMosOcTradeSyncbanksale(clt *core.SDKClient, req *mos.AlibabaMosOcTradeSyncbanksaleRequest, session string) (*mos.AlibabaMosOcTradeSyncbanksaleAPIResponse, error) {
    var resp mos.AlibabaMosOcTradeSyncbanksaleAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
