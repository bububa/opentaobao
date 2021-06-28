package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词改价 APIResponse
alibaba.scbp.ad.keyword.price.update

关键词改价
*/
type AlibabaScbpAdKeywordPriceUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_ad_keyword_price_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 修改关键词价格是否成功
    
    Result   bool `json:"result,omitempty" xml:"