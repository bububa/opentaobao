package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
推客平台订单回流 APIResponse
alibaba.trade.aliance.create

推客平台订单回流
*/
type AlibabaTradeAlianceCreateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_trade_aliance_create_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 订单创建结果
    
    Result   *AlibabaTradeAlianceCreateResultModel `json:"result,omitempty" xml:"