package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
供应商/分销商关闭采购申请/经销采购单 APIResponse
taobao.fenxiao.dealer.requisitionorder.close

供应商或分销商关闭采购申请/经销采购单
*/
type TaobaoFenxiaoDealerRequisitionorderCloseAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoFenxiaoDealerRequisitionorderCloseResponse `json:"fenxiao_dealer_requisitionorder_close_response,omitempty"` 
    TaobaoFenxiaoDealerRequisitionorderCloseResponse
}

/* model for simplify = false
type TaobaoFenxiaoDealerRequisitionorderCloseResponse struct {

    // 操作是否成功。true：成功；false：失败。
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoFenxiaoDealerRequisitionorderCloseResponse struct {

    // 操作是否成功。true：成功；false：失败。
    IsSuccess   bool `json:"is_success,omitempty"`

}
