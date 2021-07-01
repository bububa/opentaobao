package moscm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/moscm"
)

/* 
可售库存查询 
alibaba.mos.goods.inventory.getinventorys

查询商品的可售、在库和占库数量
*/
func AlibabaMosGoodsInventoryGetinventorys(clt *core.SDKClient, req *moscm.AlibabaMosGoodsInventoryGetinventorysAPIRequest, session string) (*moscm.AlibabaMosGoodsInventoryGetinventorysAPIResponse, error) {
    var resp moscm.AlibabaMosGoodsInventoryGetinventorysAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
