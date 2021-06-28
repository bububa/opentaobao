package refund

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询买家申请的退款列表 APIResponse
taobao.refunds.apply.get

查询买家申请的退款列表，且查询外店的退款列表时需要指定交易类型
*/
type TaobaoRefundsApplyGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoRefundsApplyGetResponse `json:"refunds_apply_get_response,omitempty"` 
    TaobaoRefundsApplyGetResponse
}

/* model for simplify = false
type TaobaoRefundsApplyGetResponse struct {

    // 搜索到的退款信息列表
    
    Refunds  struct {
        Refund  []Refund `json:"refund,omitempty"`
    } `json:"refunds,omitempty"`
    

    // 搜索到的交易信息总数
    
    TotalResults   int64 `json:"total_results,omitempty"`
    

}
*/

type TaobaoRefundsApplyGetResponse struct {

    // 搜索到的退款信息列表
    Refunds   []Refund `json:"refunds,omitempty"`

    // 搜索到的交易信息总数
    TotalResults   int64 `json:"total_results,omitempty"`

}
