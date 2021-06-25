package jst

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jst"
)

/* 
异常订单信息获取 
taobao.oc.tradetrace.alerts.get

提供订单预警模块的数据查询接口
*/
func TaobaoOcTradetraceAlertsGet(clt *core.SDKClient, req *jst.TaobaoOcTradetraceAlertsGetRequest, session string) (*jst.TaobaoOcTradetraceAlertsGetResponse, error) {
    var resp jst.TaobaoOcTradetraceAlertsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
