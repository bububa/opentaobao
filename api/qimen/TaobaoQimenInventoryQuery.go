package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
库存查询接口（多商品） 
taobao.qimen.inventory.query

ERP调用奇门的接口,查询商品的库存量
*/
func TaobaoQimenInventoryQuery(clt *core.SDKClient, req *qimen.TaobaoQimenInventoryQueryAPIRequest, session string) (*qimen.TaobaoQimenInventoryQueryAPIResponse, error) {
    var resp qimen.TaobaoQimenInventoryQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
