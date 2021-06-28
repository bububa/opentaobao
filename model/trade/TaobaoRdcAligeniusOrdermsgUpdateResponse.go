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
    // Response *TaobaoRdcAligeniusOrdermsgUpdateResponse `json:"rdc_aligenius_ordermsg_update_response,omitempty"` 
    TaobaoRdcAligeniusOrdermsgUpdateResponse
}

/* model for simplify = false
type TaobaoRdcAligeniusOrdermsgUpdateResponse struct {

    // result
    
    Result  *struct {
        TaobaoRdcAligeniusOrdermsgUpdateResult  *TaobaoRdcAligeniusOrdermsgUpdateResult `json:"taobao_rdc_aligenius_ordermsg_update_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoRdcAligeniusOrdermsgUpdateResponse struct {

    // result
    Result   *TaobaoRdcAligeniusOrdermsgUpdateResult `json:"result,omitempty"`

}
