package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫奖池鉴权接口 APIResponse
alibaba.interact.isvlottery.isvdraw

鉴权接口，为tida.isvdraw接口鉴权
*/
type AlibabaInteractIsvlotteryIsvdrawAPIResponse struct {
    model.CommonResponse
    AlibabaInteractIsvlotteryIsvdrawResponse
}

type AlibabaInteractIsvlotteryIsvdrawResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_isvlottery_isvdraw_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 无用参数
    
    Stub   string `json:"stub,omitempty" xml:"stub,omitempty"`

    
}
