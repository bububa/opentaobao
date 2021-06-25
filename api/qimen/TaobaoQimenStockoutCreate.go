package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
出库单创建接口 
taobao.qimen.stockout.create

ERP调用奇门接口，创建出库单信息
*/
func TaobaoQimenStockoutCreate(clt *core.SDKClient, req *qimen.TaobaoQimenStockoutCreateRequest, session string) (*qimen.TaobaoQimenStockoutCreateResponse, error) {
    var resp qimen.TaobaoQimenStockoutCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
