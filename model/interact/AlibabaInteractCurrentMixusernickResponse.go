package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
手淘混淆nick开放接口鉴权专用 APIResponse
alibaba.interact.current.mixusernick

手淘混淆nick开放接口鉴权专用，无数据输入输出。
*/
type AlibabaInteractCurrentMixusernickAPIResponse struct {
    model.CommonResponse
    AlibabaInteractCurrentMixusernickResponse
}

type AlibabaInteractCurrentMixusernickResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_current_mixusernick_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result=0
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
