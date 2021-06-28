package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
AG商家账号校验 APIResponse
taobao.rdc.aligenius.account.validate

提供应对接AG的erp系统查询其旗下的商家是否为AG商家
*/
type TaobaoRdcAligeniusAccountValidateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoRdcAligeniusAccountValidateResponse `json:"rdc_aligenius_account_validate_response,omitempty"` 
    TaobaoRdcAligeniusAccountValidateResponse
}

/* model for simplify = false
type TaobaoRdcAligeniusAccountValidateResponse struct {

    // result
    
    Result  *struct {
        TaobaoRdcAligeniusAccountValidateResult  *TaobaoRdcAligeniusAccountValidateResult `json:"taobao_rdc_aligenius_account_validate_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoRdcAligeniusAccountValidateResponse struct {

    // result
    Result   *TaobaoRdcAligeniusAccountValidateResult `json:"result,omitempty"`

}
