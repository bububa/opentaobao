package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
外贸直通车加词 APIResponse
alibaba.scbp.ad.keyword.add

外贸直通车加词服务
*/
type AlibabaScbpAdKeywordAddAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpAdKeywordAddResponse `json:"alibaba_scbp_ad_keyword_add_response,omitempty"`
}

type AlibabaScbpAdKeywordAddResponse struct {

    // 请求加入的词
    Keyword   string `json:"keyword,omitempty"`

    // 该词是否加入成功
    IsAdded   bool `json:"is_added,omitempty"`

    // 加词失败的原因
    InvalidType   string `json:"invalid_type,omitempty"`

    // 系统中存在归一化重复的词
    RepeatKeyword   string `json:"repeat_keyword,omitempty"`

}
