package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
供应商/分销商通过采购申请/经销采购单申请 APIResponse
taobao.fenxiao.dealer.requisitionorder.agree

供应商或分销商通过采购申请/经销采购单审核
*/
type TaobaoFenxiaoDealerRequisitionorderAgreeAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoFenxiaoDealerRequisitionorderAgreeResponse `json:"fenxiao_dealer_requisitionorder_agree_response,omitempty"` 
    TaobaoFenxiaoDealerRequisitionorderAgreeResponse
}

/* model for simplify = false
type TaobaoFenxiaoDealerRequisitionorderAgreeResponse struct {

    // 操作是否成功。true：成功；false：失败。
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoFenxiaoDealerRequisitionorderAgreeResponse struct {

    // 操作是否成功。true：成功；false：失败。
    IsSuccess   bool `json:"is_success,omitempty"`

}
