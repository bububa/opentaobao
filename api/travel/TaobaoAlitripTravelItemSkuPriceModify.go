package travel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/travel"
)

/* 
【API3.0】日期级别日历价格库存修改，增量维护 
taobao.alitrip.travel.item.sku.price.modify

【API3.0】日期级别日历价格库存增量维护
*/
func TaobaoAlitripTravelItemSkuPriceModify(clt *core.SDKClient, req *travel.TaobaoAlitripTravelItemSkuPriceModifyRequest, session string) (*travel.TaobaoAlitripTravelItemSkuPriceModifyResponse, error) {
    var resp travel.TaobaoAlitripTravelItemSkuPriceModifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
