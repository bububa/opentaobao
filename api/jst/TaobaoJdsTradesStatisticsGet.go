package jst

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jst"
)

/* 
获取订单数量统计结果 
taobao.jds.trades.statistics.get

获取订单数量统计结果
*/
func TaobaoJdsTradesStatisticsGet(clt *core.SDKClient, req *jst.TaobaoJdsTradesStatisticsGetRequest, session string) (*jst.TaobaoJdsTradesStatisticsGetResponse, error) {
    var resp jst.TaobaoJdsTradesStatisticsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
