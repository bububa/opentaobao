package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
设置一批关键词的信息 APIResponse
taobao.simba.keywords.pricevon.set

设置一批关键词的信息，包含无线出价、计算机出价和关键词匹配方式
*/
type TaobaoSimbaKeywordsPricevonSetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaKeywordsPricevonSetResponse `json:"taobao_simba_keywords_pricevon_set_response,omitempty"`
}

type TaobaoSimbaKeywordsPricevonSetResponse struct {

    // 成功设置关键词价格的关键词列表
    Keywords   []Keyword `json:"keywords,omitempty"`

}
