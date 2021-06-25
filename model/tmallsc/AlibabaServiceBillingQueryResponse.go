package tmallsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
服务平台结算出账信息 APIResponse
alibaba.service.billing.query

服务平台结算单明细查询服务
*/
type AlibabaServiceBillingQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaServiceBillingQueryResponse `json:"alibaba_service_billing_query_response,omitempty"`
}

type AlibabaServiceBillingQueryResponse struct {

    // 结算明细查询结果
    SettlementDetailQueryResult   *FulfilplatformResult `json:"settlement_detail_query_result,omitempty"`

}
