package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
补发单物流信息回传 APIResponse
taobao.rdc.aligenius.warehouse.resend.logistics.msg.post

补发单erp物流信息回传平台
*/
type TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIResponse struct {
    model.CommonResponse
    Response *TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostResponse `json:"taobao_rdc_aligenius_warehouse_resend_logistics_msg_post_response,omitempty"`
}

type TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostResponse struct {

    // result
    Result   *TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostResult `json:"result,omitempty"`

}
