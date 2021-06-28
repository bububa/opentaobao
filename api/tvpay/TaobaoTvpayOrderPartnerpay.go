package tvpay

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvpay"
)

/* 
tv支付第三方支付订单 
taobao.tvpay.order.partnerpay

tv支付第三方发起并支付订单（使用设备授权）
*/
func TaobaoTvpayOrderPartnerpay(clt *core.SDKClient, req *tvpay.TaobaoTvpayOrderPartnerpayRequest, session string) (*tvpay.TaobaoTvpayOrderPartnerpayAPIResponse, error) {
    var resp tvpay.TaobaoTvpayOrderPartnerpayAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
