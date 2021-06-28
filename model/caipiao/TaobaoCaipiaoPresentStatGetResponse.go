package caipiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取卖家按天统计的彩票赠送数据 APIResponse
taobao.caipiao.present.stat.get

查询卖家一段时间内按天统计的彩票赠送数据，只支持查询90天以内的数据.
*/
type TaobaoCaipiaoPresentStatGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoCaipiaoPresentStatGetResponse `json:"caipiao_present_stat_get_response,omitempty"` 
    TaobaoCaipiaoPresentStatGetResponse
}

/* model for simplify = false
type TaobaoCaipiaoPresentStatGetResponse struct {

    // 查询的结果集大小
    
    TotalResults   int64 `json:"total_results,omitempty"`
    

    // 查询的结果集
    
    Results  struct {
        LotteryWangcaiPresentStat  []LotteryWangcaiPresentStat `json:"lottery_wangcai_present_stat,omitempty"`
    } `json:"results,omitempty"`
    

}
*/

type TaobaoCaipiaoPresentStatGetResponse struct {

    // 查询的结果集大小
    TotalResults   int64 `json:"total_results,omitempty"`

    // 查询的结果集
    Results   []LotteryWangcaiPresentStat `json:"results,omitempty"`

}
