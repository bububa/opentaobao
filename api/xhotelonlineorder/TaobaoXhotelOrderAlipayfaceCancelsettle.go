package xhotelonlineorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotelonlineorder"
)

/* 
信用住订单取消结算接口 
taobao.xhotel.order.alipayface.cancelsettle

信用住订单由于客人为出现等原因，最终取消结算,一定要在结算后2个小时之内调用，否则不会成功。
*/
func TaobaoXhotelOrderAlipayfaceCancelsettle(clt *core.SDKClient, req *xhotelonlineorder.TaobaoXhotelOrderAlipayfaceCancelsettleRequest, session string) (*xhotelonlineorder.TaobaoXhotelOrderAlipayfaceCancelsettleAPIResponse, error) {
    var resp xhotelonlineorder.TaobaoXhotelOrderAlipayfaceCancelsettleAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
