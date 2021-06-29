package opentrade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/opentrade"
)

/* 
查询尖货活动信息 
taobao.opentrade.activity.query

尖货交易活动信息配置，查询尖货活动信息
*/
func TaobaoOpentradeActivityQuery(clt *core.SDKClient, req *opentrade.TaobaoOpentradeActivityQueryRequest, session string) (*opentrade.TaobaoOpentradeActivityQueryAPIResponse, error) {
    var resp opentrade.TaobaoOpentradeActivityQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
