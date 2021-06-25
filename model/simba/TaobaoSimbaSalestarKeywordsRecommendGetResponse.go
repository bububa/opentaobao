package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
销量明星api相关接口 APIResponse
taobao.simba.salestar.keywords.recommend.get

取得一个推广组的推荐关键词列表
*/
type TaobaoSimbaSalestarKeywordsRecommendGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaSalestarKeywordsRecommendGetResponse `json:"taobao_simba_salestar_keywords_recommend_get_response,omitempty"`
}

type TaobaoSimbaSalestarKeywordsRecommendGetResponse struct {

    // 推荐词分页对象，当输入的页码大于最大数值时，将返回最大的page_no值，并且结果中的数据列表为空值
    RecommendWords   *RecommendWordPage `json:"recommend_words,omitempty"`

}
