package jipiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jipiao"
)

/* 
代理商确认机票订单接口 
taobao.alitrip.seller.order.confirm

此接口用于代理商确认机票订单。
*/
func TaobaoAlitripSellerOrderConfirm(clt *core.SDKClient, req *jipiao.TaobaoAlitripSellerOrderConfirmAPIRequest, session string) (*jipiao.TaobaoAlitripSellerOrderConfirmAPIResponse, error) {
    var resp jipiao.TaobaoAlitripSellerOrderConfirmAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
