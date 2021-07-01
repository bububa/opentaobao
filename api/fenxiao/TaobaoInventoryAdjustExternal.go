package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
非交易库存调整单 
taobao.inventory.adjust.external

建议使用新接口：taobao.inventory.merchant.adjust ，该接口会逐步停用。
商家非交易调整库存，调拨出库、盘点等时调用
*/
func TaobaoInventoryAdjustExternal(clt *core.SDKClient, req *fenxiao.TaobaoInventoryAdjustExternalAPIRequest, session string) (*fenxiao.TaobaoInventoryAdjustExternalAPIResponse, error) {
    var resp fenxiao.TaobaoInventoryAdjustExternalAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
