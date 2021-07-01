package aedropshiper

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aedropshiper"
)

/* 
买家查询订单详情 
aliexpress.trade.ds.order.get

买家查询订单详情，用于dropshipper
*/
func AliexpressTradeDsOrderGet(clt *core.SDKClient, req *aedropshiper.AliexpressTradeDsOrderGetAPIRequest, session string) (*aedropshiper.AliexpressTradeDsOrderGetAPIResponse, error) {
    var resp aedropshiper.AliexpressTradeDsOrderGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
