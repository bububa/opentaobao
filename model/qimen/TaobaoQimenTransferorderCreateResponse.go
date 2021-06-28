package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
调拨单创建 APIResponse
taobao.qimen.transferorder.create

调拨单创建
*/
type TaobaoQimenTransferorderCreateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenTransferorderCreateResponse `json:"qimen_transferorder_create_response,omitempty"` 
    TaobaoQimenTransferorderCreateResponse
}

/* model for simplify = false
type TaobaoQimenTransferorderCreateResponse struct {

    // 
    
    Response  *struct {
        TaobaoQimenTransferorderCreateStruct  *TaobaoQimenTransferorderCreateStruct `json:"taobao_qimen_transferorder_create_struct,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenTransferorderCreateResponse struct {

    // 
    Response   *TaobaoQimenTransferorderCreateStruct `json:"response,omitempty"`

}
