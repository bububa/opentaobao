package traderate

import (
    "github.com/bububa/opentaobao/model"
)

/* 
评价大家印象印象短语接口 APIResponse
taobao.traderate.impr.imprwords.get

根据淘宝后台类目的一级类目和叶子类目
*/
type TaobaoTraderateImprImprwordsGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTraderateImprImprwordsGetResponse `json:"traderate_impr_imprwords_get_response,omitempty"` 
    TaobaoTraderateImprImprwordsGetResponse
}

/* model for simplify = false
type TaobaoTraderateImprImprwordsGetResponse struct {

    // 返回类目下所有大家印象的标签
    
    ImprWords  struct {
        String  []string `json:"string,omitempty"`
    } `json:"impr_words,omitempty"`
    

}
*/

type TaobaoTraderateImprImprwordsGetResponse struct {

    // 返回类目下所有大家印象的标签
    ImprWords   []string `json:"impr_words,omitempty"`

}
