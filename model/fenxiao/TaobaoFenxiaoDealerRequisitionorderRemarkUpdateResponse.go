package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
修改经销采购单备注 APIResponse
taobao.fenxiao.dealer.requisitionorder.remark.update

供应商修改经销采购单备注
*/
type TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoFenxiaoDealerRequisitionorderRemarkUpdateResponse `json:"fenxiao_dealer_requisitionorder_remark_update_response,omitempty"` 
    TaobaoFenxiaoDealerRequisitionorderRemarkUpdateResponse
}

/* model for simplify = false
type TaobaoFenxiaoDealerRequisitionorderRemarkUpdateResponse struct {

    // 操作是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoFenxiaoDealerRequisitionorderRemarkUpdateResponse struct {

    // 操作是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
