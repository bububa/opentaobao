package trade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/trade"
)

/* 
获取单笔交易的部分信息(商家应用使用) 
taobao.open.trade.get

获取单笔交易的部分信息</br>
1.入参fields中传入buyer_nick ，才能返回buyer_open_id
*/
func TaobaoOpenTradeGet(clt *core.SDKClient, req *trade.TaobaoOpenTradeGetAPIRequest, session string) (*trade.TaobaoOpenTradeGetAPIResponse, error) {
    var resp trade.TaobaoOpenTradeGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
