package travel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/travel"
)

/* 
【API3.0】商品级别日历价格库存修改，全量覆盖 
taobao.alitrip.travel.item.sku.override

旅行度假新商品日历价格库存信息修改接口 第三版。提供商家通过TOP API方式修改商品sku信息。
*/
func TaobaoAlitripTravelItemSkuOverride(clt *core.SDKClient, req *travel.TaobaoAlitripTravelItemSkuOverrideRequest, session string) (*travel.TaobaoAlitripTravelItemSkuOverrideResponse, error) {
    var resp travel.TaobaoAlitripTravelItemSkuOverrideAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
