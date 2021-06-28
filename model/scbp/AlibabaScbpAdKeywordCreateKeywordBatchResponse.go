package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
关键词添加 APIResponse
alibaba.scbp.ad.keyword.create.keyword.batch

关键词添加
*/
type AlibabaScbpAdKeywordCreateKeywordBatchAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpAdKeywordCreateKeywordBatchResponse `json:"alibaba_scbp_ad_keyword_create_keyword_batch_response,omitempty"` 
    AlibabaScbpAdKeywordCreateKeywordBatchResponse
}

/* model for simplify = false
type AlibabaScbpAdKeywordCreateKeywordBatchResponse struct {

    // 返回错误集合
    
    ResultList  struct {
        ErrorKeyword  []ErrorKeyword `json:"error_keyword,omitempty"`
    } `json:"result_list,omitempty"`
    

}
*/

type AlibabaScbpAdKeywordCreateKeywordBatchResponse struct {

    // 返回错误集合
    ResultList   []ErrorKeyword `json:"result_list,omitempty"`

}
