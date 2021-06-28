package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词批量改价 APIResponse
alibaba.scbp.ad.keyword.price.batchupdate

关键词批量改价
*/
type AlibabaScbpAdKeywordPriceBatchupdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_ad_keyword_price_batchupdate_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 修改失败关键词列表
    
    KeywordErrorResultList   []KeywordErrorResultDto `json:"keyword_error_result_list,omitempty" xml:"