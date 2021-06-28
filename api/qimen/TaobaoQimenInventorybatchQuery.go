package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
商品单仓批次库存查询接口 
taobao.qimen.inventorybatch.query

ERP 通过该接口查询指定商品的单仓批次库存
*/
func TaobaoQimenInventorybatchQuery(clt *core.SDKClient, req *qimen.TaobaoQimenInventorybatchQueryRequest, session string) (*qimen.TaobaoQimenInventorybatchQueryAPIResponse, error) {
    var resp qimen.TaobaoQimenInventorybatchQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
