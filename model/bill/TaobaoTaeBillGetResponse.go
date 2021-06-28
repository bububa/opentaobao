package bill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
tae查询单笔账单明细 APIResponse
taobao.tae.bill.get

查询单笔账单明细
*/
type TaobaoTaeBillGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTaeBillGetResponse `json:"tae_bill_get_response,omitempty"` 
    TaobaoTaeBillGetResponse
}

/* model for simplify = false
type TaobaoTaeBillGetResponse struct {

    // 账单明细
    
    Bill  *struct {
        BillDto  *BillDto `json:"bill_dto,omitempty"`
    } `json:"bill,omitempty"`
    

}
*/

type TaobaoTaeBillGetResponse struct {

    // 账单明细
    Bill   *BillDto `json:"bill,omitempty"`

}
