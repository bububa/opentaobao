package tvpay

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvpay"
)

/* 
tv支付授权查询 
taobao.tvpay.auth.query

查询该用户在指定设备上是否有支付授权
*/
func TaobaoTvpayAuthQuery(clt *core.SDKClient, req *tvpay.TaobaoTvpayAuthQueryAPIRequest, session string) (*tvpay.TaobaoTvpayAuthQueryAPIResponse, error) {
    var resp tvpay.TaobaoTvpayAuthQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
