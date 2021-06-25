package qianniu

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qianniu"
)

/* 
销售订单总额统计 
taobao.daogoubao.order.statistics.total

对接千牛端数字中心
*/
func TaobaoDaogoubaoOrderStatisticsTotal(clt *core.SDKClient, req *qianniu.TaobaoDaogoubaoOrderStatisticsTotalRequest, session string) (*qianniu.TaobaoDaogoubaoOrderStatisticsTotalResponse, error) {
    var resp qianniu.TaobaoDaogoubaoOrderStatisticsTotalAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
