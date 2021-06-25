package refund

import (
    "github.com/bububa/opentaobao/model"
)

/* 
取消发货 APIResponse
taobao.rdc.aligenius.sendgoods.cancel

提供商家在仅退款中发送取消发货状态
*/
type TaobaoRdcAligeniusSendgoodsCancelAPIResponse struct {
    model.CommonResponse
    Response *TaobaoRdcAligeniusSendgoodsCancelResponse `json:"taobao_rdc_aligenius_sendgoods_cancel_response,omitempty"`
}

type TaobaoRdcAligeniusSendgoodsCancelResponse struct {

    // result
    Result   *TaobaoRdcAligeniusSendgoodsCancelResult `json:"result,omitempty"`

}
