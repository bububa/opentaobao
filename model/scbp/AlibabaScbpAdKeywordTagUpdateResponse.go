package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
修改关键词所属分组 APIResponse
alibaba.scbp.ad.keyword.tag.update

修改关键词所属分组
*/
type AlibabaScbpAdKeywordTagUpdateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpAdKeywordTagUpdateResponse `json:"alibaba_scbp_ad_keyword_tag_update_response,omitempty"` 
    AlibabaScbpAdKeywordTagUpdateResponse
}

/* model for simplify = false
type AlibabaScbpAdKeywordTagUpdateResponse struct {

    // 实际修改的关键词数
    
    Result   int64 `json:"result,omitempty"`
    

}
*/

type AlibabaScbpAdKeywordTagUpdateResponse struct {

    // 实际修改的关键词数
    Result   int64 `json:"result,omitempty"`

}
