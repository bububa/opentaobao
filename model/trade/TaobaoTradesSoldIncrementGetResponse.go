package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询卖家已卖出的增量交易数据（根据修改时间） APIResponse
taobao.trades.sold.increment.get

搜索当前会话用户作为卖家已卖出的增量交易数据（只能获取到三个月以内的交易信息）
<br/>1. 一次请求只能查询时间跨度为一天的增量交易记录，即end_modified - start_modified <= 1天。
<br/>2. 返回的数据结果是以订单的修改时间倒序排列的，通过从后往前翻页的方式可以避免漏单问题。
<br/>3. 返回的数据结果只包含了订单的部分数据，可通过taobao.trade.fullinfo.get获取订单详情。
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=tradeapi" target="_blank">点击查看更多交易API说明</a></strong>
*/
type TaobaoTradesSoldIncrementGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTradesSoldIncrementGetResponse `json:"trades_sold_increment_get_response,omitempty"` 
    TaobaoTradesSoldIncrementGetResponse
}

/* model for simplify = false
type TaobaoTradesSoldIncrementGetResponse struct {

    // 搜索到的交易信息总数
    
    TotalResults   int64 `json:"total_results,omitempty"`
    

    // 搜索到的交易信息列表，返回的Trade和Order中包含的具体信息为入参fields请求的字段信息
    
    Trades  struct {
        Trade  []Trade `json:"trade,omitempty"`
    } `json:"trades,omitempty"`
    

    // 是否存在下一页
    
    HasNext   bool `json:"has_next,omitempty"`
    

}
*/

type TaobaoTradesSoldIncrementGetResponse struct {

    // 搜索到的交易信息总数
    TotalResults   int64 `json:"total_results,omitempty"`

    // 搜索到的交易信息列表，返回的Trade和Order中包含的具体信息为入参fields请求的字段信息
    Trades   []Trade `json:"trades,omitempty"`

    // 是否存在下一页
    HasNext   bool `json:"has_next,omitempty"`

}
