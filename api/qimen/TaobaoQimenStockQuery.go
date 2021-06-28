package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
库存查询接口（多条件） 
taobao.qimen.stock.query

ERP调用奇门的接口,查询商品的库存量
*/
func TaobaoQimenStockQuery(clt *core.SDKClient, req *qimen.TaobaoQimenStockQueryRequest, session string) (*qimen.TaobaoQimenStockQueryAPIResponse, error) {
    var resp qimen.TaobaoQimenStockQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
