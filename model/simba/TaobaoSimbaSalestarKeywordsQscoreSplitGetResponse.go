package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
(新)销量明星质量分相关接口 APIResponse
taobao.simba.salestar.keywords.qscore.split.get

获取关键词新的质量分
*/
type TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaSalestarKeywordsQscoreSplitGetResponse `json:"taobao_simba_salestar_keywords_qscore_split_get_response,omitempty"`
}

type TaobaoSimbaSalestarKeywordsQscoreSplitGetResponse struct {

    // result
    Result   *TaobaoSimbaSalestarKeywordsQscoreSplitGetResultDto `json:"result,omitempty"`

}