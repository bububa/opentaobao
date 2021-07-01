package jipiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jipiao"
)

/* 
【机票代理商】确认退票 
taobao.alitrip.seller.refund.confirmreturn

确认退票
*/
func TaobaoAlitripSellerRefundConfirmreturn(clt *core.SDKClient, req *jipiao.TaobaoAlitripSellerRefundConfirmreturnAPIRequest, session string) (*jipiao.TaobaoAlitripSellerRefundConfirmreturnAPIResponse, error) {
    var resp jipiao.TaobaoAlitripSellerRefundConfirmreturnAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
