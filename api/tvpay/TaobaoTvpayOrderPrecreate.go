package tvpay

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvpay"
)

/* 
tv支付预下单 
taobao.tvpay.order.precreate

tv支付预下单
*/
func TaobaoTvpayOrderPrecreate(clt *core.SDKClient, req *tvpay.TaobaoTvpayOrderPrecreateAPIRequest, session string) (*tvpay.TaobaoTvpayOrderPrecreateAPIResponse, error) {
    var resp tvpay.TaobaoTvpayOrderPrecreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
