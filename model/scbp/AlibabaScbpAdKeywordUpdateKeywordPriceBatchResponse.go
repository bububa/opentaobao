package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改关键词价格 API返回值 
alibaba.scbp.ad.keyword.update.keyword.price.batch

修改关键词价格
*/
type AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdKeywordUpdateKeywordPriceBatchResponse
}

// 修改关键词价格 成功返回结果
type AlibabaScbpAdKeywordUpdateKeywordPriceBatchResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_update_keyword_price_batch_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 错误信息集合
    ResultList   []ErrorKeyword `json:"result_list,omitempty" xml:"result_list>error_keyword,omitempty"`
}
