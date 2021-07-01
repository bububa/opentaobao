package dutyfree

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/dutyfree"
)

/* 
对外库存查询接口 
alibaba.dutyfree.stock.query

对外部服务提供库存查询接口
*/
func AlibabaDutyfreeStockQuery(clt *core.SDKClient, req *dutyfree.AlibabaDutyfreeStockQueryAPIRequest, session string) (*dutyfree.AlibabaDutyfreeStockQueryAPIResponse, error) {
    var resp dutyfree.AlibabaDutyfreeStockQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
