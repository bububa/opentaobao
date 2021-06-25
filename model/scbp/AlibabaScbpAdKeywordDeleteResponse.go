package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
外贸直通车删除关键词 APIResponse
alibaba.scbp.ad.keyword.delete

外贸直通车删除关键词
*/
type AlibabaScbpAdKeywordDeleteAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpAdKeywordDeleteResponse `json:"alibaba_scbp_ad_keyword_delete_response,omitempty"`
}

type AlibabaScbpAdKeywordDeleteResponse struct {

    // 删除关键词是否成功
    Result   bool `json:"result,omitempty"`

}
