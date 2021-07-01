package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
无线端抽奖接口 API返回值 
alibaba.interact.lotterydraw.dodraw

商家抽奖平台无线端抽奖接口开放
*/
type AlibabaInteractLotterydrawDodrawAPIResponse struct {
    model.CommonResponse
    AlibabaInteractLotterydrawDodrawAPIResponseModel
}

// 无线端抽奖接口 成功返回结果
type AlibabaInteractLotterydrawDodrawAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_interact_lotterydraw_dodraw_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AlibabaInteractLotterydrawDodrawResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
