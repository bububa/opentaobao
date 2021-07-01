package axintrade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/axintrade"
)

/* 
修改资金单接口 
taobao.alitrip.axin.trans.fund.update

阿信供销平台-修改资金单接口
*/
func TaobaoAlitripAxinTransFundUpdate(clt *core.SDKClient, req *axintrade.TaobaoAlitripAxinTransFundUpdateAPIRequest, session string) (*axintrade.TaobaoAlitripAxinTransFundUpdateAPIResponse, error) {
    var resp axintrade.TaobaoAlitripAxinTransFundUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
