package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取词的相关词 APIResponse
taobao.simba.insight.relatedwords.get

获取给定词的若干相关词，返回结果中越相关的权重越大，排在越前面，根据number参数对返回结果进行截断。
*/
type TaobaoSimbaInsightRelatedwordsGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaInsightRelatedwordsGetResponse `json:"taobao_simba_insight_relatedwords_get_response,omitempty"`
}

type TaobaoSimbaInsightRelatedwordsGetResponse struct {

    // 相关词列表，最多可传100个。
    RelatedWordsResultList   []InsightRelatedWords `json:"related_words_result_list,omitempty"`

}
