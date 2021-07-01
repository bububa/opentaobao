package jipiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jipiao"
)

/* 
【机票代理商订单】确认退款 
taobao.alitrip.seller.refundmoney.confirm

代理人确认退票申请单的退款
*/
func TaobaoAlitripSellerRefundmoneyConfirm(clt *core.SDKClient, req *jipiao.TaobaoAlitripSellerRefundmoneyConfirmAPIRequest, session string) (*jipiao.TaobaoAlitripSellerRefundmoneyConfirmAPIResponse, error) {
    var resp jipiao.TaobaoAlitripSellerRefundmoneyConfirmAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
