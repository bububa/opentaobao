package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
订单全链路状态统计差异比较 APIResponse
taobao.jds.trades.statistics.diff

订单全链路状态统计差异比较
*/
type TaobaoJdsTradesStatisticsDiffAPIResponse struct {
    model.CommonResponse
    Response *TaobaoJdsTradesStatisticsDiffResponse `json:"taobao_jds_trades_statistics_diff_response,omitempty"`
}

type TaobaoJdsTradesStatisticsDiffResponse struct {

    // pre_status比post_status多的tid列表
    Tids   []Number `json:"tids,omitempty"`

    // 总记录数
    TotalResults   int64 `json:"total_results,omitempty"`

}
