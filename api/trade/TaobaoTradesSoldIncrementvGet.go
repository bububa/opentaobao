package trade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/trade"
)

/* 
查询卖家已卖出的增量交易数据（根据入库时间） 
taobao.trades.sold.incrementv.get

搜索当前会话用户作为卖家已卖出的增量交易数据（只能获取到三个月以内的交易信息） 
<br/>1. 一次请求只能查询时间跨度为一天的增量交易记录，即end_create - start_create <= 1天。 
<br/>2. 返回的数据结果是以订单入库时间的倒序排列的(该时间和订单修改时间不同)，通过从后往前翻页的方式可以避免漏单问题。 
<br/>3. 返回的数据结果只包含了订单的部分数据，可通过taobao.trade.fullinfo.get获取订单详情。 
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=tradeapi" target="_blank">点击查看更多交易API说明</a></strong>
*/
func TaobaoTradesSoldIncrementvGet(clt *core.SDKClient, req *trade.TaobaoTradesSoldIncrementvGetRequest, session string) (*trade.TaobaoTradesSoldIncrementvGetResponse, error) {
    var resp trade.TaobaoTradesSoldIncrementvGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
