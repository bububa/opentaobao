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
    Response *TaobaoJdsTradesStatisticsGetResponse `json:"taobao_jds_trades_statistics_get_response,omitempty"`
}

type TaobaoJdsTradesStatisticsGetResponse struct {

    // 订单状态计数值
    StatusInfos   []TradeStat `json:"status_infos,omitempty"`

}
