package tanx

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tanx"
)

/* 
tanx竞价失败反馈api 
taobao.tanx.biddingrefuses.get

竞价失败反馈根据创意id查询API提供
*/
func TaobaoTanxBiddingrefusesGet(clt *core.SDKClient, req *tanx.TaobaoTanxBiddingrefusesGetRequest, session string) (*tanx.TaobaoTanxBiddingrefusesGetResponse, error) {
    var resp tanx.TaobaoTanxBiddingrefusesGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
