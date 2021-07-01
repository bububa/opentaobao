package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
商品优惠详情查询 
taobao.ump.promotion.get

商品优惠详情查询，可查询商品设置的详细优惠。包括限时折扣，满就送等官方优惠以及第三方优惠。
*/
func TaobaoUmpPromotionGet(clt *core.SDKClient, req *product.TaobaoUmpPromotionGetAPIRequest, session string) (*product.TaobaoUmpPromotionGetAPIResponse, error) {
    var resp product.TaobaoUmpPromotionGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
