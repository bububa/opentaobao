package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
外贸直通车加词 APIResponse
alibaba.scbp.ad.keyword.add

外贸直通车加词服务
*/
type AlibabaScbpAdKeywordAddAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdKeywordAddResponse
}

type AlibabaScbpAdKeywordAddResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_add_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 请求加入的词
    
    Keyword   string `json:"keyword,omitempty" xml:"keyword,omitempty"`

    
    // 该词是否加入成功
    
    IsAdded   bool `json:"is_added,omitempty" xml:"is_added,omitempty"`

    
    // 加词失败的原因
    
    InvalidType   string `json:"invalid_type,omitempty" xml:"invalid_type,omitempty"`

    
    // 系统中存在归一化重复的词
    
    RepeatKeyword   string `json:"repeat_keyword,omitempty" xml:"repeat_keyword,omitempty"`

    
}
