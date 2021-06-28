package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
退款查询 APIResponse
alibaba.mos.onsite.trade.queryrefund

商户可使用该接口查询退款请求是否执行成功。
*/
type AlibabaMosOnsiteTradeQueryrefundAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_mos_onsite_trade_queryrefund_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlibabaMosOnsiteTradeQueryrefundResultDo `json:"result,omitempty" xml:"