package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
出库单确认接口 APIResponse
taobao.qimen.stockout.confirm

货品出库后，WMS将状态回传给ERP
*/
type TaobaoQimenStockoutConfirmAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQimenStockoutConfirmResponse `json:"taobao_qimen_stockout_confirm_response,omitempty"`
}

type TaobaoQimenStockoutConfirmResponse struct {

    // 
    Response   *TaobaoQimenStockoutConfirmStruct `json:"response,omitempty"`

}
