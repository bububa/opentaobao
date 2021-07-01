package lstwarehouse

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/lstwarehouse"
)

/* 
零售通经销商商品库存设置 
alibaba.lst.ic.stock.items.update

零售通经销商商品库存设置
*/
func AlibabaLstIcStockItemsUpdate(clt *core.SDKClient, req *lstwarehouse.AlibabaLstIcStockItemsUpdateAPIRequest, session string) (*lstwarehouse.AlibabaLstIcStockItemsUpdateAPIResponse, error) {
    var resp lstwarehouse.AlibabaLstIcStockItemsUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
