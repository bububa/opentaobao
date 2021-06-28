package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
网厅垂直信息查询接口 APIResponse
taobao.trade.wtvertical.get

网厅订单垂直信息的查询
*/
type TaobaoTradeWtverticalGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTradeWtverticalGetResponse `json:"trade_wtvertical_get_response,omitempty"` 
    TaobaoTradeWtverticalGetResponse
}

/* model for simplify = false
type TaobaoTradeWtverticalGetResponse struct {

    // 返回交易垂直信息的数据结构
    
    WtextResults  struct {
        WtExtResult  []WtExtResult `json:"wt_ext_result,omitempty"`
    } `json:"wtext_results,omitempty"`
    

    // 返回查询记录的条数
    
    TotalResults   int64 `json:"total_results,omitempty"`
    

}
*/

type TaobaoTradeWtverticalGetResponse struct {

    // 返回交易垂直信息的数据结构
    WtextResults   []WtExtResult `json:"wtext_results,omitempty"`

    // 返回查询记录的条数
    TotalResults   int64 `json:"total_results,omitempty"`

}
