package nlp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
文本语言词法分析 APIResponse
taobao.nlp.word

提供文本语言处理中的词法分析功能,开放中文分词和词权重计算功能。
*/
type TaobaoNlpWordAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoNlpWordResponse `json:"nlp_word_response,omitempty"` 
    TaobaoNlpWordResponse
}

/* model for simplify = false
type TaobaoNlpWordResponse struct {

    // 返回词法分析的结果
    
    Wordresult  *struct {
        WordResult  *WordResult `json:"word_result,omitempty"`
    } `json:"wordresult,omitempty"`
    

}
*/

type TaobaoNlpWordResponse struct {

    // 返回词法分析的结果
    Wordresult   *WordResult `json:"wordresult,omitempty"`

}
