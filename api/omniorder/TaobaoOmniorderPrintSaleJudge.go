package omniorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/omniorder"
)

/* 
导购员判断 
taobao.omniorder.print.sale.judge

用于判断当前子账号是否导购员
*/
func TaobaoOmniorderPrintSaleJudge(clt *core.SDKClient, req *omniorder.TaobaoOmniorderPrintSaleJudgeAPIRequest, session string) (*omniorder.TaobaoOmniorderPrintSaleJudgeAPIResponse, error) {
    var resp omniorder.TaobaoOmniorderPrintSaleJudgeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
