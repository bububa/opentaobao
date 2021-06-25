package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
新质量分服务 APIResponse
taobao.simba.keywords.qscore.split.get

获取关键词新的质量分
*/
type TaobaoSimbaKeywordsQscoreSplitGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaKeywordsQscoreSplitGetResponse `json:"taobao_simba_keywords_qscore_split_get_response,omitempty"`
}

type TaobaoSimbaKeywordsQscoreSplitGetResponse struct {

    // result
    Result   *TaobaoSimbaKeywordsQscoreSplitGetResultDto `json:"result,omitempty"`

}
