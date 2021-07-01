package tanx

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tanx"
)

/* 
对外部dsp提供交易id查询接口 
taobao.tanx.deal.get

对外部dsp提供交易id查询接口
*/
func TaobaoTanxDealGet(clt *core.SDKClient, req *tanx.TaobaoTanxDealGetAPIRequest, session string) (*tanx.TaobaoTanxDealGetAPIResponse, error) {
    var resp tanx.TaobaoTanxDealGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
