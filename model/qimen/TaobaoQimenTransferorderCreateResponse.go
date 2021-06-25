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
    Response *TaobaoQimenTransferorderCreateResponse `json:"taobao_qimen_transferorder_create_response,omitempty"`
}

type TaobaoQimenTransferorderCreateResponse struct {

    // 
    Response   *TaobaoQimenTransferorderCreateStruct `json:"response,omitempty"`

}
