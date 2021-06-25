package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
查询后端商品仓库库存 
tmall.inventory.query.forstore

商家查询后端商品仓库库存
*/
func TmallInventoryQueryForstore(clt *core.SDKClient, req *fenxiao.TmallInventoryQueryForstoreRequest, session string) (*fenxiao.TmallInventoryQueryForstoreResponse, error) {
    var resp fenxiao.TmallInventoryQueryForstoreAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
