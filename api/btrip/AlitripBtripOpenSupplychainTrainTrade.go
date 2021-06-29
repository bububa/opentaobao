package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
商旅火车票交易流水接口 
alitrip.btrip.open.supplychain.train.trade

商旅火车票交易流水接口
*/
func AlitripBtripOpenSupplychainTrainTrade(clt *core.SDKClient, req *btrip.AlitripBtripOpenSupplychainTrainTradeRequest, session string) (*btrip.AlitripBtripOpenSupplychainTrainTradeAPIResponse, error) {
    var resp btrip.AlitripBtripOpenSupplychainTrainTradeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
