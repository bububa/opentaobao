package category

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询两个渠道之间的固定映射关系，不通过算法兜底 APIResponse
alibaba.imap.fixedmapping.query

查询两个渠道之间的固定映射关系，不通过算法兜底
*/
type AlibabaImapFixedmappingQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaImapFixedmappingQueryResponse `json:"alibaba_imap_fixedmapping_query_response,omitempty"`
}

type AlibabaImapFixedmappingQueryResponse struct {

    // 接口返回model
    Result   *AlibabaImapFixedmappingQueryResult `json:"result,omitempty"`

}
