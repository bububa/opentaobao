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
    AlibabaInteractLotterydrawDodrawResponse
}

type AlibabaInteractLotterydrawDodrawResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_lotterydraw_dodraw_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlibabaInteractLotterydrawDodrawResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
