package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
批量启动暂停推广词状态 APIResponse
alibaba.scbp.ad.keyword.status.batchupdate

批量启动暂停关键词推广状态
*/
type AlibabaScbpAdKeywordStatusBatchupdateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpAdKeywordStatusBatchupdateResponse `json:"alibaba_scbp_ad_keyword_status_batchupdate_response,omitempty"`
}

type AlibabaScbpAdKeywordStatusBatchupdateResponse struct {

    // 修改失败关键词列表
    KeywordErrorResultList   []KeywordErrorResultDto `json:"keyword_error_result_list,omitempty"`

}
