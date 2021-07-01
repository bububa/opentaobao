package util

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/util"
)

/* 
通过订单获取对应买家的openUID 
taobao.openuid.get.bytrade

通过订单获取对应买家的openUID,需要卖家授权
*/
func TaobaoOpenuidGetBytrade(clt *core.SDKClient, req *util.TaobaoOpenuidGetBytradeAPIRequest, session string) (*util.TaobaoOpenuidGetBytradeAPIResponse, error) {
    var resp util.TaobaoOpenuidGetBytradeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
