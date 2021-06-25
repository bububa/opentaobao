package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
订单消息状态回传 APIResponse
taobao.rdc.aligenius.ordermsg.update

用于订单消息处理状态回传
*/
type TaobaoRdcAligeniusOrdermsgUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoRdcAligeniusOrdermsgUpdateResponse `json:"taobao_rdc_aligenius_ordermsg_update_response,omitempty"`
}

type TaobaoRdcAligeniusOrdermsgUpdateResponse struct {

    // result
    Result   *TaobaoRdcAligeniusOrdermsgUpdateResult `json:"result,omitempty"`

}
