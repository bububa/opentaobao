package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取订单数量统计结果 APIResponse
taobao.jds.trades.statistics.get

获取订单数量统计结果
*/
type TaobaoJdsTradesStatisticsGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoJdsTradesStatisticsGetResponse `json:"jds_trades_statistics_get_response,omitempty"` 
    TaobaoJdsTradesStatisticsGetResponse
}

/* model for simplify = false
type TaobaoJdsTradesStatisticsGetResponse struct {

    // 订单状态计数值
    
    StatusInfos  struct {
        TradeStat  []TradeStat `json:"trade_stat,omitempty"`
    } `json:"status_infos,omitempty"`
    

}
*/

type TaobaoJdsTradesStatisticsGetResponse struct {

    // 订单状态计数值
    StatusInfos   []TradeStat `json:"status_infos,omitempty"`

}
