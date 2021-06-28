package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
互动到店抽奖 APIResponse
alibaba.interact.isvlottery.idraw

互动到店抽奖
*/
type AlibabaInteractIsvlotteryIdrawAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_interact_isvlottery_idraw_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 抽奖中奖信息
    
    Data   *LotteryProxyResult `json:"data,omitempty" xml:"