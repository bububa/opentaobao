package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
数据关联关系查询 APIResponse
alibaba.wdk.marketing.open.data.relation.query

数据关联关系查询
*/
type AlibabaWdkMarketingOpenDataRelationQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_marketing_open_data_relation_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果信息
    
    Result   *WdkMarketOpenResult `json:"result,omitempty" xml:"