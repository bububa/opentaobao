package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
外贸直通车批量删除关键词 APIResponse
alibaba.scbp.ad.keyword.batchdelete

外贸直通车批量删除关键词
*/
type AlibabaScbpAdKeywordBatchdeleteAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpAdKeywordBatchdeleteResponse `json:"alibaba_scbp_ad_keyword_batchdelete_response,omitempty"`
}

type AlibabaScbpAdKeywordBatchdeleteResponse struct {

    // 删除关键词是否成功
    Result   bool `json:"result,omitempty"`

}
