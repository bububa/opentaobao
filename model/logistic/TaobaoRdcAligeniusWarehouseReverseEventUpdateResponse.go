package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
销退单事件回传接口 APIResponse
taobao.rdc.aligenius.warehouse.reverse.event.update

用于erp回传销退单相关信息到平台
*/
type TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"rdc_aligenius_warehouse_reverse_event_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *TaobaoRdcAligeniusWarehouseReverseEventUpdateResult `json:"result,omitempty" xml:"