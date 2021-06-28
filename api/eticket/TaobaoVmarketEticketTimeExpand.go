package eticket

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/eticket"
)

/* 
订单延时接口 
taobao.vmarket.eticket.time.expand

提供码商操作订单延期接口
*/
func TaobaoVmarketEticketTimeExpand(clt *core.SDKClient, req *eticket.TaobaoVmarketEticketTimeExpandRequest, session string) (*eticket.TaobaoVmarketEticketTimeExpandAPIResponse, error) {
    var resp eticket.TaobaoVmarketEticketTimeExpandAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
