package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
线下新退款接口（专为老退款接口调用） APIResponse
alibaba.mos.onsite.trade.oldrefund

线下新退款接口（专为老退款接口调用）。新接口支付支付宝、手淘、天猫，老接口退款时，需要调用该接口退新单，并适配老接口响应参数返回
*/
type AlibabaMosOnsiteTradeOldrefundAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_mos_onsite_trade_oldrefund_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 交易退款响应
    
    Result   string `json:"result,omitempty" xml:"