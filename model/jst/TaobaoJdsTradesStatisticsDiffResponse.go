package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单全链路状态统计差异比较 APIResponse
taobao.jds.trades.statistics.diff

订单全链路状态统计差异比较
*/
type TaobaoJdsTradesStatisticsDiffAPIResponse struct {
    model.CommonResponse
    TaobaoJdsTradesStatisticsDiffResponse
}

type TaobaoJdsTradesStatisticsDiffResponse struct {
    XMLName xml.Name `xml:"jds_trades_statistics_diff_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // pre_status比post_status多的tid列表
    
    Tids   []int64 `json:"tids,omitempty" xml:"tids>int64,omitempty"`
    
    
    // 总记录数
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`

    
}
