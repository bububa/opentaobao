package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除关键词 APIResponse
alibaba.scbp.ad.keyword.delete.keyword.batch

删除关键词
*/
type AlibabaScbpAdKeywordDeleteKeywordBatchAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpAdKeywordDeleteKeywordBatchResponse `json:"alibaba_scbp_ad_keyword_delete_keyword_batch_response,omitempty"`
}

type AlibabaScbpAdKeywordDeleteKeywordBatchResponse struct {

    // 返回结果
    Result   int64 `json:"result,omitempty"`

}
