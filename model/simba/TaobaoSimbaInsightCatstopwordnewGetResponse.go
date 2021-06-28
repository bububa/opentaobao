package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取类目下最热门的词 APIResponse
taobao.simba.insight.catstopwordnew.get

按照某个维度，查询某个类目下最热门的词，维度有点击，展现，花费，点击率等，具体可以按哪些字段进行排序，参考参数说明，比如选择了impression，则返回该类目下展现量最高那几个词。
*/
type TaobaoSimbaInsightCatstopwordnewGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaInsightCatstopwordnewGetResponse `json:"simba_insight_catstopwordnew_get_response,omitempty"` 
    TaobaoSimbaInsightCatstopwordnewGetResponse
}

/* model for simplify = false
type TaobaoSimbaInsightCatstopwordnewGetResponse struct {

    // 类目下热门词详细数据
    
    TopwordDataList  struct {
        InsightWordDataUnderCatDTO  []InsightWordDataUnderCatDTO `json:"insight_word_data_under_cat_dto,omitempty"`
    } `json:"topword_data_list,omitempty"`
    

}
*/

type TaobaoSimbaInsightCatstopwordnewGetResponse struct {

    // 类目下热门词详细数据
    TopwordDataList   []InsightWordDataUnderCatDTO `json:"topword_data_list,omitempty"`

}
