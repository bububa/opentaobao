package mtopopen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
回传抽奖相关参数 APIResponse
alibaba.interact.lotteryactivity.register

提供接口供三方应用将数据回传到平台
*/
type AlibabaInteractLotteryactivityRegisterAPIResponse struct {
    model.CommonResponse
    AlibabaInteractLotteryactivityRegisterResponse
}

type AlibabaInteractLotteryactivityRegisterResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_lotteryactivity_register_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaInteractLotteryactivityRegisterResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
