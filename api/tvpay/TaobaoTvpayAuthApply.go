package tvpay

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvpay"
)

/* 
tv支付申请设备授权 
taobao.tvpay.auth.apply

为用户在指定设备上申请支付授权
*/
func TaobaoTvpayAuthApply(clt *core.SDKClient, req *tvpay.TaobaoTvpayAuthApplyAPIRequest, session string) (*tvpay.TaobaoTvpayAuthApplyAPIResponse, error) {
    var resp tvpay.TaobaoTvpayAuthApplyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
