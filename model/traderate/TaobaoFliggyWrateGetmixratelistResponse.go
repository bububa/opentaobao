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
    Response *TaobaoFliggyWrateGetmixratelistResponse `json:"taobao_fliggy_wrate_getmixratelist_response,omitempty"`
}

type TaobaoFliggyWrateGetmixratelistResponse struct {

    // 接口返回model
    Result   *TaobaoFliggyWrateGetmixratelistResult `json:"result,omitempty"`

}
