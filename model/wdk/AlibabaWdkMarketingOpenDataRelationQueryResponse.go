package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
数据关联关系查询 APIResponse
alibaba.wdk.marketing.open.data.relation.query

数据关联关系查询
*/
type AlibabaWdkMarketingOpenDataRelationQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMarketingOpenDataRelationQueryResponse `json:"alibaba_wdk_marketing_open_data_relation_query_response,omitempty"`
}

type AlibabaWdkMarketingOpenDataRelationQueryResponse struct {

    // 结果信息
    Result   *WdkMarketOpenResult `json:"result,omitempty"`

}
