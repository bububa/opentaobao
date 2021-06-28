package jae

import (
    "github.com/bububa/opentaobao/model"
)

/* 
活动平台弱登录接口 APIResponse
taobao.aplatform.weakget

无线活动平台的开放接口，提供商品信息等的读操作
*/
type TaobaoAplatformWeakgetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAplatformWeakgetResponse `json:"aplatform_weakget_response,omitempty"` 
    TaobaoAplatformWeakgetResponse
}

/* model for simplify = false
type TaobaoAplatformWeakgetResponse struct {

    // result
    
    Result  *struct {
        TaobaoAplatformWeakgetResult  *TaobaoAplatformWeakgetResult `json:"taobao_aplatform_weakget_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoAplatformWeakgetResponse struct {

    // result
    Result   *TaobaoAplatformWeakgetResult `json:"result,omitempty"`

}
