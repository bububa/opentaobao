package moscm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/moscm"
)

/* 
以盘点方式调整库存：传入商品实际库存 
alibaba.mos.goods.synchinventorybycounting

以盘点方式调整库存：传入商品实际库存
盘点单自动判断数量增减
*/
func AlibabaMosGoodsSynchinventorybycounting(clt *core.SDKClient, req *moscm.AlibabaMosGoodsSynchinventorybycountingAPIRequest, session string) (*moscm.AlibabaMosGoodsSynchinventorybycountingAPIResponse, error) {
    var resp moscm.AlibabaMosGoodsSynchinventorybycountingAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
