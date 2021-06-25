package wlb

import (
    "github.com/bububa/opentaobao/model"
)

/* 
BMS出库通知 APIResponse
cainiao.bms.order.consign.confirm

BMS出库后，通知ISV
*/
type CainiaoBmsOrderConsignConfirmAPIResponse struct {
    model.CommonResponse
    Response *CainiaoBmsOrderConsignConfirmResponse `json:"cainiao_bms_order_consign_confirm_response,omitempty"`
}

type CainiaoBmsOrderConsignConfirmResponse struct {

    // 返回值
    Result   *ResultDO `json:"result,omitempty"`

}
