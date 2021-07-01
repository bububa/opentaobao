package tbk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tbk"
)

/* 
sku维度最优优惠券信息 
taobao.tbk.sku.best.coupon

根据itemid和skuid查询最优优惠券信息
*/
func TaobaoTbkSkuBestCoupon(clt *core.SDKClient, req *tbk.TaobaoTbkSkuBestCouponAPIRequest, session string) (*tbk.TaobaoTbkSkuBestCouponAPIResponse, error) {
    var resp tbk.TaobaoTbkSkuBestCouponAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
