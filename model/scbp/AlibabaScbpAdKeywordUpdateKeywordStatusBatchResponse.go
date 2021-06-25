package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
修改关键词状态 APIResponse
alibaba.scbp.ad.keyword.update.keyword.status.batch

修改关键词状态
*/
type AlibabaScbpAdKeywordUpdateKeywordStatusBatchAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpAdKeywordUpdateKeywordStatusBatchResponse `json:"alibaba_scbp_ad_keyword_update_keyword_status_batch_response,omitempty"`
}

type AlibabaScbpAdKeywordUpdateKeywordStatusBatchResponse struct {

    // 返回错误集合
    ResultList   []ErrorKeyword `json:"result_list,omitempty"`

}
