package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
关键词批量改价 APIResponse
alibaba.scbp.ad.keyword.price.batchupdate

关键词批量改价
*/
type AlibabaScbpAdKeywordPriceBatchupdateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpAdKeywordPriceBatchupdateResponse `json:"alibaba_scbp_ad_keyword_price_batchupdate_response,omitempty"` 
    AlibabaScbpAdKeywordPriceBatchupdateResponse
}

/* model for simplify = false
type AlibabaScbpAdKeywordPriceBatchupdateResponse struct {

    // 修改失败关键词列表
    
    KeywordErrorResultList  struct {
        KeywordErrorResultDto  []KeywordErrorResultDto `json:"keyword_error_result_dto,omitempty"`
    } `json:"keyword_error_result_list,omitempty"`
    

}
*/

type AlibabaScbpAdKeywordPriceBatchupdateResponse struct {

    // 修改失败关键词列表
    KeywordErrorResultList   []KeywordErrorResultDto `json:"keyword_error_result_list,omitempty"`

}
