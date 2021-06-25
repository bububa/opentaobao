package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
库存预占取消接口 APIResponse
taobao.qimen.inventoryreserve.cancel

库存预占取消
*/
type TaobaoQimenInventoryreserveCancelAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQimenInventoryreserveCancelResponse `json:"taobao_qimen_inventoryreserve_cancel_response,omitempty"`
}

type TaobaoQimenInventoryreserveCancelResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}
