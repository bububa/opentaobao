package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
双11到店互动无线端抽奖接口鉴权 APIResponse
alibaba.interact.wireless.draw

双11到店互动无线端mtop接口开放鉴权接口，无入参出参，无安全风险，mtop接口开发 坯子
*/
type AlibabaInteractWirelessDrawAPIResponse struct {
    model.CommonResponse
    AlibabaInteractWirelessDrawResponse
}

type AlibabaInteractWirelessDrawResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_wireless_draw_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
