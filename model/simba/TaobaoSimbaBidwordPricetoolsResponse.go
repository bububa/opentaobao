package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
关键词出价指导工具（新） APIResponse
taobao.simba.bidword.pricetools

关键词出价指导工具（新）
*/
type TaobaoSimbaBidwordPricetoolsAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaBidwordPricetoolsResponse `json:"taobao_simba_bidword_pricetools_response,omitempty"`
}

type TaobaoSimbaBidwordPricetoolsResponse struct {

    // true 表示符合准入，false不符合
    ResultList   *PriceSuggestionDto `json:"result_list,omitempty"`

}
