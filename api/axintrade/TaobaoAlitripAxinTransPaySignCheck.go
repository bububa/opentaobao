package axintrade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/axintrade"
)

/* 
阿信支付宝验签服务 
taobao.alitrip.axin.trans.pay.sign.check

阿信支付宝验签服务
*/
func TaobaoAlitripAxinTransPaySignCheck(clt *core.SDKClient, req *axintrade.TaobaoAlitripAxinTransPaySignCheckRequest, session string) (*axintrade.TaobaoAlitripAxinTransPaySignCheckAPIResponse, error) {
    var resp axintrade.TaobaoAlitripAxinTransPaySignCheckAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
