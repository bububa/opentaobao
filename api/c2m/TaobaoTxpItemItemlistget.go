package c2m

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/c2m"
)

/* 
淘小铺商品接口 
taobao.txp.item.itemlistget

淘小铺商品的查询服务。
*/
func TaobaoTxpItemItemlistget(clt *core.SDKClient, req *c2m.TaobaoTxpItemItemlistgetRequest, session string) (*c2m.TaobaoTxpItemItemlistgetAPIResponse, error) {
    var resp c2m.TaobaoTxpItemItemlistgetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
