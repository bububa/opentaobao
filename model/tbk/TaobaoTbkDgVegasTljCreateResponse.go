package tbk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-淘礼金创建 APIResponse
taobao.tbk.dg.vegas.tlj.create

创建淘礼金
*/
type TaobaoTbkDgVegasTljCreateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTbkDgVegasTljCreateResponse `json:"tbk_dg_vegas_tlj_create_response,omitempty"` 
    TaobaoTbkDgVegasTljCreateResponse
}

/* model for simplify = false
type TaobaoTbkDgVegasTljCreateResponse struct {

    // result
    
    Result  *struct {
        TaobaoTbkDgVegasTljCreateResult  *TaobaoTbkDgVegasTljCreateResult `json:"taobao_tbk_dg_vegas_tlj_create_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoTbkDgVegasTljCreateResponse struct {

    // result
    Result   *TaobaoTbkDgVegasTljCreateResult `json:"result,omitempty"`

}
