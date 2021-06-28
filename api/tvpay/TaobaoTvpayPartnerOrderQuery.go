package tvpay

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvpay"
)

/* 
商户查询订单 
taobao.tvpay.partner.order.query

给商户提供的查询订单状态的API
*/
func TaobaoTvpayPartnerOrderQuery(clt *core.SDKClient, req *tvpay.TaobaoTvpayPartnerOrderQueryRequest, session string) (*tvpay.TaobaoTvpayPartnerOrderQueryAPIResponse, error) {
    var resp tvpay.TaobaoTvpayPartnerOrderQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
