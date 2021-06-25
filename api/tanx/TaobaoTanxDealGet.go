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
func TaobaoTanxDealGet(clt *core.SDKClient, req *tanx.TaobaoTanxDealGetRequest, session string) (*tanx.TaobaoTanxDealGetResponse, error) {
    var resp tanx.TaobaoTanxDealGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
