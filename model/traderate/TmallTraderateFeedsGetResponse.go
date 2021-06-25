package traderate

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询子订单对应的评价、追评以及语义标签 APIResponse
tmall.traderate.feeds.get

通过子订单ID获取天猫订单对应的评价，追评，以及对应的语义标签
*/
type TmallTraderateFeedsGetAPIResponse struct {
    model.CommonResponse
    Response *TmallTraderateFeedsGetResponse `json:"tmall_traderate_feeds_get_response,omitempty"`
}

type TmallTraderateFeedsGetResponse struct {

    // 返回评价信息
    TmallRateInfo   *TmallRateInfo `json:"tmall_rate_info,omitempty"`

}
