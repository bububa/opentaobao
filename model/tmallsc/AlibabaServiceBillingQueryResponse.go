package tmallsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务平台结算出账信息 APIResponse
alibaba.service.billing.query

服务平台结算单明细查询服务
*/
type AlibabaServiceBillingQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_service_billing_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结算明细查询结果
    
    SettlementDetailQueryResult   *FulfilplatformResult `json:"settlement_detail_query_result,omitempty" xml:"