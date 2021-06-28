package traderate

import (
    "github.com/bububa/opentaobao/model"
)

/* 
搜索评价信息 APIResponse
taobao.traderates.get

搜索评价信息，只能获取距今180天内的评价记录(只支持查询卖家给出或得到的评价)。
*/
type TaobaoTraderatesGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTraderatesGetResponse `json:"traderates_get_response,omitempty"` 
    TaobaoTraderatesGetResponse
}

/* model for simplify = false
type TaobaoTraderatesGetResponse struct {

    // 评价列表。返回的TradeRate包含的具体信息为入参fields请求的字段信息
    
    TradeRates  struct {
        TradeRate  []TradeRate `json:"trade_rate,omitempty"`
    } `json:"trade_rates,omitempty"`
    

    // 搜索到的评价总数。相同的查询时间段条件下，最大只能获取总共1500条评价记录。所以当评价大于等于1500时 ISV可以通过start_date及end_date来进行拆分，以保证可以查询到全部数据
    
    TotalResults   int64 `json:"total_results,omitempty"`
    

    // 当使用use_has_next时返回信息，如果还有下一页则返回true
    
    HasNext   bool `json:"has_next,omitempty"`
    

}
*/

type TaobaoTraderatesGetResponse struct {

    // 评价列表。返回的TradeRate包含的具体信息为入参fields请求的字段信息
    TradeRates   []TradeRate `json:"trade_rates,omitempty"`

    // 搜索到的评价总数。相同的查询时间段条件下，最大只能获取总共1500条评价记录。所以当评价大于等于1500时 ISV可以通过start_date及end_date来进行拆分，以保证可以查询到全部数据
    TotalResults   int64 `json:"total_results,omitempty"`

    // 当使用use_has_next时返回信息，如果还有下一页则返回true
    HasNext   bool `json:"has_next,omitempty"`

}
