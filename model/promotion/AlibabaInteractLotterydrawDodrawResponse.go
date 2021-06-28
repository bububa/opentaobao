package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
无线端抽奖接口 APIResponse
alibaba.interact.lotterydraw.dodraw

商家抽奖平台无线端抽奖接口开放
*/
type AlibabaInteractLotterydrawDodrawAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_interact_lotterydraw_dodraw_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlibabaInteractLotterydrawDodrawResultDto `json:"result,omitempty" xml:"