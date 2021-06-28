package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
修改关键词价格 APIResponse
alibaba.scbp.ad.keyword.update.keyword.price.batch

修改关键词价格
*/
type AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpAdKeywordUpdateKeywordPriceBatchResponse `json:"alibaba_scbp_ad_keyword_update_keyword_price_batch_response,omitempty"` 
    AlibabaScbpAdKeywordUpdateKeywordPriceBatchResponse
}

/* model for simplify = false
type AlibabaScbpAdKeywordUpdateKeywordPriceBatchResponse struct {

    // 错误信息集合
    
    ResultList  struct {
        ErrorKeyword  []ErrorKeyword `json:"error_keyword,omitempty"`
    } `json:"result_list,omitempty"`
    

}
*/

type AlibabaScbpAdKeywordUpdateKeywordPriceBatchResponse struct {

    // 错误信息集合
    ResultList   []ErrorKeyword `json:"result_list,omitempty"`

}
