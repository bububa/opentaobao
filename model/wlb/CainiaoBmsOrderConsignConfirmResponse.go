package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
BMS出库通知 APIResponse
cainiao.bms.order.consign.confirm

BMS出库后，通知ISV
*/
type CainiaoBmsOrderConsignConfirmAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"cainiao_bms_order_consign_confirm_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回值
    
    Result   *ResultDO `json:"result,omitempty" xml:"