package jst

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jst"
)

/* 
订单全链路状态统计差异比较 
taobao.jds.trades.statistics.diff

订单全链路状态统计差异比较
*/
func TaobaoJdsTradesStatisticsDiff(clt *core.SDKClient, req *jst.TaobaoJdsTradesStatisticsDiffRequest, session string) (*jst.TaobaoJdsTradesStatisticsDiffResponse, error) {
    var resp jst.TaobaoJdsTradesStatisticsDiffAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
