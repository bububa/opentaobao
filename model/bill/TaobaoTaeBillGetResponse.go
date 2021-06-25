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
    Response *TaobaoTaeBillGetResponse `json:"taobao_tae_bill_get_response,omitempty"`
}

type TaobaoTaeBillGetResponse struct {

    // 账单明细
    Bill   *BillDto `json:"bill,omitempty"`

}
