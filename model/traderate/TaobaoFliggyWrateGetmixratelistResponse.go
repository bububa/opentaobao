package traderate

import (
    "github.com/bububa/opentaobao/model"
)

/* 
飞猪通用评价接口 APIResponse
taobao.fliggy.wrate.getmixratelist

飞猪评价通用接口
*/
type TaobaoFliggyWrateGetmixratelistAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoFliggyWrateGetmixratelistResponse `json:"fliggy_wrate_getmixratelist_response,omitempty"` 
    TaobaoFliggyWrateGetmixratelistResponse
}

/* model for simplify = false
type TaobaoFliggyWrateGetmixratelistResponse struct {

    // 接口返回model
    
    Result  *struct {
        TaobaoFliggyWrateGetmixratelistResult  *TaobaoFliggyWrateGetmixratelistResult `json:"taobao_fliggy_wrate_getmixratelist_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoFliggyWrateGetmixratelistResponse struct {

    // 接口返回model
    Result   *TaobaoFliggyWrateGetmixratelistResult `json:"result,omitempty"`

}
