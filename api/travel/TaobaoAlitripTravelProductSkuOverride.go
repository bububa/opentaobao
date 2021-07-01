package travel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/travel"
)

/* 
（供销）产品级别日历价格库存修改，全量覆盖 
taobao.alitrip.travel.product.sku.override

（供销）产品级别日历价格库存修改，全量覆盖
*/
func TaobaoAlitripTravelProductSkuOverride(clt *core.SDKClient, req *travel.TaobaoAlitripTravelProductSkuOverrideAPIRequest, session string) (*travel.TaobaoAlitripTravelProductSkuOverrideAPIResponse, error) {
    var resp travel.TaobaoAlitripTravelProductSkuOverrideAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
