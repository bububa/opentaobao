package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
预购结算数据统计 APIResponse
alibaba.mj.presale.settlement.statistics

预购结算数据统计
*/
type AlibabaMjPresaleSettlementStatisticsAPIResponse struct {
    model.CommonResponse
    AlibabaMjPresaleSettlementStatisticsResponse
}

type AlibabaMjPresaleSettlementStatisticsResponse struct {
    XMLName xml.Name `xml:"alibaba_mj_presale_settlement_statistics_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // data
    
    Data   *AlibabaMjPresaleSettlementStatisticsData `json:"data,omitempty" xml:"data,omitempty"`

    
}
