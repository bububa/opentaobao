package wlb

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wlb"
)

/* 
商家库存调整 
cainiao.merchant.inventory.adjust

商家仓库存调整接口，目前仅支持全量更新
*/
func CainiaoMerchantInventoryAdjust(clt *core.SDKClient, req *wlb.CainiaoMerchantInventoryAdjustAPIRequest, session string) (*wlb.CainiaoMerchantInventoryAdjustAPIResponse, error) {
    var resp wlb.CainiaoMerchantInventoryAdjustAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
