package refund

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/refund"
)

/* 
卖家回填物流信息 
taobao.rp.returngoods.refill

卖家收到货物回填物流信息，如果买家已经回填物流信息，则接口报错，目前仅支持天猫订单。
*/
func TaobaoRpReturngoodsRefill(clt *core.SDKClient, req *refund.TaobaoRpReturngoodsRefillAPIRequest, session string) (*refund.TaobaoRpReturngoodsRefillAPIResponse, error) {
    var resp refund.TaobaoRpReturngoodsRefillAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
