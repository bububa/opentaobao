package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取订单数量统计结果 APIResponse
taobao.jds.trades.statistics.get

获取订单数量统计结果
*/
type TaobaoJdsTradesStatisticsGetAPIResponse struct {
    model.CommonResponse
    TaobaoJdsTradesStatisticsGetResponse
}

type TaobaoJdsTradesStatisticsGetResponse struct {
    XMLName xml.Name `xml:"jds_trades_statistics_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 订单状态计数值
    
    StatusInfos   []TradeStat `json:"status_infos,omitempty" xml:"status_infos>trade_stat,omitempty"`
    
    
}
