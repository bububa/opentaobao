package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
入库单创建接口 APIResponse
taobao.qimen.entryorder.create

ERP调用接口，创建入库单;
*/
type TaobaoQimenEntryorderCreateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQimenEntryorderCreateResponse `json:"taobao_qimen_entryorder_create_response,omitempty"`
}

type TaobaoQimenEntryorderCreateResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}