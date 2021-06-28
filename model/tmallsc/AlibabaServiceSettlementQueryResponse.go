package tmallsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务平台结算单明细查询服务 APIResponse
alibaba.service.settlement.query

给服务商提供结算单明细查询功能
*/
type AlibabaServiceSettlementQueryAPIResponse struct {
    model.CommonResponse
    AlibabaServiceSettlementQueryResponse
}

type AlibabaServiceSettlementQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_service_settlement_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结算明细查询结果
    
    SettlementDetailQueryResult   *FulfilplatformResult `json:"settlement_detail_query_result,omitempty" xml:"settlement_detail_query_result,omitempty"`

    
}
