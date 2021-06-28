package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
无线端抽奖接口 APIResponse
alibaba.interact.lotterydraw.dodraw

商家抽奖平台无线端抽奖接口开放
*/
type AlibabaInteractLotterydrawDodrawAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaInteractLotterydrawDodrawResponse `json:"alibaba_interact_lotterydraw_dodraw_response,omitempty"` 
    AlibabaInteractLotterydrawDodrawResponse
}

/* model for simplify = false
type AlibabaInteractLotterydrawDodrawResponse struct {

    // result
    
    Result  *struct {
        AlibabaInteractLotterydrawDodrawResultDto  *AlibabaInteractLotterydrawDodrawResultDto `json:"alibaba_interact_lotterydraw_dodraw_result_dto,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaInteractLotterydrawDodrawResponse struct {

    // result
    Result   *AlibabaInteractLotterydrawDodrawResultDto `json:"result,omitempty"`

}
