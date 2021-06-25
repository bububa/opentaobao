package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
入库单确认接口 APIResponse
taobao.qimen.entryorder.confirm

WMS调用接口，回传入库单信息;
*/
type TaobaoQimenEntryorderConfirmAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQimenEntryorderConfirmResponse `json:"taobao_qimen_entryorder_confirm_response,omitempty"`
}

type TaobaoQimenEntryorderConfirmResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}
