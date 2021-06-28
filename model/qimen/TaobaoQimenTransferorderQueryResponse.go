package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
调拨单查询 APIResponse
taobao.qimen.transferorder.query

调拨单查询
*/
type TaobaoQimenTransferorderQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenTransferorderQueryResponse `json:"qimen_transferorder_query_response,omitempty"` 
    TaobaoQimenTransferorderQueryResponse
}

/* model for simplify = false
type TaobaoQimenTransferorderQueryResponse struct {

    // 
    
    Response  *struct {
        TaobaoQimenTransferorderQueryStruct  *TaobaoQimenTransferorderQueryStruct `json:"taobao_qimen_transferorder_query_struct,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenTransferorderQueryResponse struct {

    // 
    Response   *TaobaoQimenTransferorderQueryStruct `json:"response,omitempty"`

}
