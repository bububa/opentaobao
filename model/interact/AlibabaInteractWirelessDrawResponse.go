package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
双11到店互动无线端抽奖接口鉴权 APIResponse
alibaba.interact.wireless.draw

双11到店互动无线端mtop接口开放鉴权接口，无入参出参，无安全风险，mtop接口开发 坯子
*/
type AlibabaInteractWirelessDrawAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaInteractWirelessDrawResponse `json:"alibaba_interact_wireless_draw_response,omitempty"` 
    AlibabaInteractWirelessDrawResponse
}

/* model for simplify = false
type AlibabaInteractWirelessDrawResponse struct {

    // result
    
    Result   string `json:"result,omitempty"`
    

}
*/

type AlibabaInteractWirelessDrawResponse struct {

    // result
    Result   string `json:"result,omitempty"`

}
