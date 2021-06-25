package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
取得一个推广组的所有关键词的质量得分或者根据关键词Id列表取得一组关键词的质量得分 APIResponse
taobao.simba.keywords.qscore.get

取得一个推广组的所有关键词的质量得分列表
*/
type TaobaoSimbaKeywordsQscoreGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaKeywordsQscoreGetResponse `json:"taobao_simba_keywords_qscore_get_response,omitempty"`
}

type TaobaoSimbaKeywordsQscoreGetResponse struct {

    // 取得的关键词质量得分列表
    KeywordQscoreList   []KeywordQscore `json:"keyword_qscore_list,omitempty"`

}
