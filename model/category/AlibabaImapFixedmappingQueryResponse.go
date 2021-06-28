package category

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询两个渠道之间的固定映射关系，不通过算法兜底 APIResponse
alibaba.imap.fixedmapping.query

查询两个渠道之间的固定映射关系，不通过算法兜底
*/
type AlibabaImapFixedmappingQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_imap_fixedmapping_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *AlibabaImapFixedmappingQueryResult `json:"result,omitempty" xml:"