package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
天猫奖池鉴权接口 APIResponse
alibaba.interact.isvlottery.isvdraw

鉴权接口，为tida.isvdraw接口鉴权
*/
type AlibabaInteractIsvlotteryIsvdrawAPIResponse struct {
    model.CommonResponse
    Response *AlibabaInteractIsvlotteryIsvdrawResponse `json:"alibaba_interact_isvlottery_isvdraw_response,omitempty"`
}

type AlibabaInteractIsvlotteryIsvdrawResponse struct {

    // 无用参数
    Stub   string `json:"stub,omitempty"`

}
