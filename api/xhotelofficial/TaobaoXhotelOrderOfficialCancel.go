package xhotelofficial

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotelofficial"
)

/* 
官网信用住订单取消 
taobao.xhotel.order.official.cancel

官网信用住订单取消
*/
func TaobaoXhotelOrderOfficialCancel(clt *core.SDKClient, req *xhotelofficial.TaobaoXhotelOrderOfficialCancelAPIRequest, session string) (*xhotelofficial.TaobaoXhotelOrderOfficialCancelAPIResponse, error) {
    var resp xhotelofficial.TaobaoXhotelOrderOfficialCancelAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
