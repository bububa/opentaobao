package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
关键词改价 APIResponse
alibaba.scbp.ad.keyword.price.update

关键词改价
*/
type AlibabaScbpAdKeywordPriceUpdateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpAdKeywordPriceUpdateResponse `json:"alibaba_scbp_ad_keyword_price_update_response,omitempty"` 
    AlibabaScbpAdKeywordPriceUpdateResponse
}

/* model for simplify = false
type AlibabaScbpAdKeywordPriceUpdateResponse struct {

    // 修改关键词价格是否成功
    
    Result   bool `json:"result,omitempty"`
    

}
*/

type AlibabaScbpAdKeywordPriceUpdateResponse struct {

    // 修改关键词价格是否成功
    Result   bool `json:"result,omitempty"`

}
