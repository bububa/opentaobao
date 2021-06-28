package tmallsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
服务平台结算单明细查询服务 APIResponse
alibaba.service.settlement.query

给服务商提供结算单明细查询功能
*/
type AlibabaServiceSettlementQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaServiceSettlementQueryResponse `json:"alibaba_service_settlement_query_response,omitempty"` 
    AlibabaServiceSettlementQueryResponse
}

/* model for simplify = false
type AlibabaServiceSettlementQueryResponse struct {

    // 结算明细查询结果
    
    SettlementDetailQueryResult  *struct {
        FulfilplatformResult  *FulfilplatformResult `json:"fulfilplatform_result,omitempty"`
    } `json:"settlement_detail_query_result,omitempty"`
    

}
*/

type AlibabaServiceSettlementQueryResponse struct {

    // 结算明细查询结果
    SettlementDetailQueryResult   *FulfilplatformResult `json:"settlement_detail_query_result,omitempty"`

}
