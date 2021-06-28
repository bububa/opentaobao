package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
云闪付、银行卡销售数据回传接口 APIResponse
alibaba.mos.oc.trade.syncbanksale

云闪付、银行卡销售数据回传
*/
type AlibabaMosOcTradeSyncbanksaleAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_mos_oc_trade_syncbanksale_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // resultDTO
    
    Result   *AlibabaMosOcTradeSyncbanksaleResultDo `json:"result,omitempty" xml:"