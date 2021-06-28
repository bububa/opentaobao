package tanx

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tanx"
)

/* 
批量获取交易列表 
taobao.tanx.deals.get

批量获取交易信息
*/
func TaobaoTanxDealsGet(clt *core.SDKClient, req *tanx.TaobaoTanxDealsGetRequest, session string) (*tanx.TaobaoTanxDealsGetAPIResponse, error) {
    var resp tanx.TaobaoTanxDealsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
