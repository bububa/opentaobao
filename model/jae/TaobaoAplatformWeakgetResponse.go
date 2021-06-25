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
    Response *TaobaoAplatformWeakgetResponse `json:"taobao_aplatform_weakget_response,omitempty"`
}

type TaobaoAplatformWeakgetResponse struct {

    // result
    Result   *TaobaoAplatformWeakgetResult `json:"result,omitempty"`

}
