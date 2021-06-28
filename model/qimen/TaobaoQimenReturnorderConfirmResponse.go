package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
退货入库单确认接口 APIResponse
taobao.qimen.returnorder.confirm

taobao.qimen.returnorder.confirm
*/
type TaobaoQimenReturnorderConfirmAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenReturnorderConfirmResponse `json:"qimen_returnorder_confirm_response,omitempty"` 
    TaobaoQimenReturnorderConfirmResponse
}

/* model for simplify = false
type TaobaoQimenReturnorderConfirmResponse struct {

    // 
    
    Response  *struct {
        Response  *Response `json:"response,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenReturnorderConfirmResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}
