package elife

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/elife"
)

/* 
查询交易结果 
taobao.elife.lifecard.query

卖家在交易状态不明的情况下, 查询交易结果.
*/
func TaobaoElifeLifecardQuery(clt *core.SDKClient, req *elife.TaobaoElifeLifecardQueryRequest, session string) (*elife.TaobaoElifeLifecardQueryResponse, error) {
    var resp elife.TaobaoElifeLifecardQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
