package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
手淘混淆nick开放接口鉴权专用 APIResponse
alibaba.interact.current.mixusernick

手淘混淆nick开放接口鉴权专用，无数据输入输出。
*/
type AlibabaInteractCurrentMixusernickAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaInteractCurrentMixusernickResponse `json:"alibaba_interact_current_mixusernick_response,omitempty"` 
    AlibabaInteractCurrentMixusernickResponse
}

/* model for simplify = false
type AlibabaInteractCurrentMixusernickResponse struct {

    // result=0
    
    Result   string `json:"result,omitempty"`
    

}
*/

type AlibabaInteractCurrentMixusernickResponse struct {

    // result=0
    Result   string `json:"result,omitempty"`

}
