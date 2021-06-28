package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
轻pos品牌营销关单接口 APIResponse
alibaba.wdk.pos.trade.close

轻pos品牌营销场景，提供关单接口给外部商家
*/
type AlibabaWdkPosTradeCloseAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_pos_trade_close_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 关单结果
    
    Result   *FastBuyPosCloseResult `json:"result,omitempty" xml:"