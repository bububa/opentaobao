package travel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/travel"
)

/* 
【API3.0】套餐级别日历价格库存增删操作 
taobao.alitrip.travel.item.sku.package.modify

【API3.0】套餐级别日历价格库存增删操作
*/
func TaobaoAlitripTravelItemSkuPackageModify(clt *core.SDKClient, req *travel.TaobaoAlitripTravelItemSkuPackageModifyAPIRequest, session string) (*travel.TaobaoAlitripTravelItemSkuPackageModifyAPIResponse, error) {
    var resp travel.TaobaoAlitripTravelItemSkuPackageModifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
