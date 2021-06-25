package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
云闪付、银行卡销售数据回传接口 APIResponse
alibaba.mos.oc.trade.syncbanksale

云闪付、银行卡销售数据回传
*/
type AlibabaMosOcTradeSyncbanksaleAPIResponse struct {
    model.CommonResponse
    Response *AlibabaMosOcTradeSyncbanksaleResponse `json:"alibaba_mos_oc_trade_syncbanksale_response,omitempty"`
}

type AlibabaMosOcTradeSyncbanksaleResponse struct {

    // resultDTO
    Result   *AlibabaMosOcTradeSyncbanksaleResultDo `json:"result,omitempty"`

}
