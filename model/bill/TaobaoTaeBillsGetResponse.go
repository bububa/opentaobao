package bill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
tae查询账单明细 APIResponse
taobao.tae.bills.get

tae查询账单明细
*/
type TaobaoTaeBillsGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTaeBillsGetResponse `json:"tae_bills_get_response,omitempty"` 
    TaobaoTaeBillsGetResponse
}

/* model for simplify = false
type TaobaoTaeBillsGetResponse struct {

    // 当前页查询返回的结果数(0-100)。相同的查询时间段条件下，最大只能获取总共5000条记录。所以当大于等于5000时 ISV可以通过start_time及end_time来进行拆分，以保证可以查询到全部数据
    
    TotalResults   int64 `json:"total_results,omitempty"`
    

    // 是否存在下一页
    
    HasNext   bool `json:"has_next,omitempty"`
    

    // 账单列表
    
    Bills  struct {
        BillDto  []BillDto `json:"bill_dto,omitempty"`
    } `json:"bills,omitempty"`
    

}
*/

type TaobaoTaeBillsGetResponse struct {

    // 当前页查询返回的结果数(0-100)。相同的查询时间段条件下，最大只能获取总共5000条记录。所以当大于等于5000时 ISV可以通过start_time及end_time来进行拆分，以保证可以查询到全部数据
    TotalResults   int64 `json:"total_results,omitempty"`

    // 是否存在下一页
    HasNext   bool `json:"has_next,omitempty"`

    // 账单列表
    Bills   []BillDto `json:"bills,omitempty"`

}
