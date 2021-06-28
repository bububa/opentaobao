package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
补发单物流信息回传 APIResponse
taobao.rdc.aligenius.warehouse.resend.logistics.msg.post

补发单erp物流信息回传平台
*/
type TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"rdc_aligenius_warehouse_resend_logistics_msg_post_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostResult `json:"result,omitempty" xml:"