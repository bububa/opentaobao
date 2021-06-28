package servicecenter

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/servicecenter"
)

/* 
服务市场内购服务下单接口 
taobao.fuwu.purchase.order.confirm

通过传入服务市场商品的itemcode等信息，返回给服务商内购服务的下单链接
*/
func TaobaoFuwuPurchaseOrderConfirm(clt *core.SDKClient, req *servicecenter.TaobaoFuwuPurchaseOrderConfirmRequest, session string) (*servicecenter.TaobaoFuwuPurchaseOrderConfirmAPIResponse, error) {
    var resp servicecenter.TaobaoFuwuPurchaseOrderConfirmAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
