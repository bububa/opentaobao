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
    Response *TaobaoTbkDgVegasTljCreateResponse `json:"taobao_tbk_dg_vegas_tlj_create_response,omitempty"`
}

type TaobaoTbkDgVegasTljCreateResponse struct {

    // result
    Result   *TaobaoTbkDgVegasTljCreateResult `json:"result,omitempty"`

}
