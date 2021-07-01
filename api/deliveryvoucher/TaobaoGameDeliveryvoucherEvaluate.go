package deliveryvoucher

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/deliveryvoucher"
)

/* 
卡券评价回传 
taobao.game.deliveryvoucher.evaluate

卡券ISV回传商品评价
*/
func TaobaoGameDeliveryvoucherEvaluate(clt *core.SDKClient, req *deliveryvoucher.TaobaoGameDeliveryvoucherEvaluateAPIRequest, session string) (*deliveryvoucher.TaobaoGameDeliveryvoucherEvaluateAPIResponse, error) {
    var resp deliveryvoucher.TaobaoGameDeliveryvoucherEvaluateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
