package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
预购结算数据统计 APIResponse
alibaba.mj.presale.settlement.statistics

预购结算数据统计
*/
type AlibabaMjPresaleSettlementStatisticsAPIResponse struct {
    model.CommonResponse
    Response *AlibabaMjPresaleSettlementStatisticsResponse `json:"alibaba_mj_presale_settlement_statistics_response,omitempty"`
}

type AlibabaMjPresaleSettlementStatisticsResponse struct {

    // data
    Data   *AlibabaMjPresaleSettlementStatisticsData `json:"data,omitempty"`

}
