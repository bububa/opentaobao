package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分佣结果查询 APIResponse
alibaba.retail.commission.result.query

查询导购分佣记录
*/
type AlibabaRetailCommissionResultQueryAPIResponse struct {
    model.CommonResponse
    AlibabaRetailCommissionResultQueryResponse
}

type AlibabaRetailCommissionResultQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_retail_commission_result_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回包装类
    
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
