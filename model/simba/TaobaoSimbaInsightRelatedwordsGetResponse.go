package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取词的相关词 API返回值 
taobao.simba.insight.relatedwords.get

获取给定词的若干相关词，返回结果中越相关的权重越大，排在越前面，根据number参数对返回结果进行截断。
*/
type TaobaoSimbaInsightRelatedwordsGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaInsightRelatedwordsGetResponse
}

// 获取词的相关词 成功返回结果
type TaobaoSimbaInsightRelatedwordsGetResponse struct {
    XMLName xml.Name `xml:"simba_insight_relatedwords_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 相关词列表，最多可传100个。
    RelatedWordsResultList   []InsightRelatedWords `json:"related_words_result_list,omitempty" xml:"related_words_result_list>insight_related_words,omitempty"`
}
