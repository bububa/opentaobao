package axintrade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/axintrade"
)

/* 
通过外部订单ID查询所有资金单 
taobao.alitrip.axin.trans.fund.query.by.order

阿信供销平台-通过外部订单ID查询所有资金单
*/
func TaobaoAlitripAxinTransFundQueryByOrder(clt *core.SDKClient, req *axintrade.TaobaoAlitripAxinTransFundQueryByOrderAPIRequest, session string) (*axintrade.TaobaoAlitripAxinTransFundQueryByOrderAPIResponse, error) {
    var resp axintrade.TaobaoAlitripAxinTransFundQueryByOrderAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
