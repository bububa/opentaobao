package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
产品库存修改 
taobao.fenxiao.product.quantity.update

修改产品库存信息，支持全量修改以及增量修改两种方式
*/
func TaobaoFenxiaoProductQuantityUpdate(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductQuantityUpdateAPIRequest, session string) (*fenxiao.TaobaoFenxiaoProductQuantityUpdateAPIResponse, error) {
    var resp fenxiao.TaobaoFenxiaoProductQuantityUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
