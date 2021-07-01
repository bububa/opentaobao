package jipiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jipiao"
)

/* 
【机票代理商】拒绝退票 
taobao.alitrip.seller.refund.refusereturn

拒绝退票
*/
func TaobaoAlitripSellerRefundRefusereturn(clt *core.SDKClient, req *jipiao.TaobaoAlitripSellerRefundRefusereturnAPIRequest, session string) (*jipiao.TaobaoAlitripSellerRefundRefusereturnAPIResponse, error) {
    var resp jipiao.TaobaoAlitripSellerRefundRefusereturnAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
